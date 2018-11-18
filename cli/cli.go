package cli

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/andygrunwald/go-jira"
	"github.com/spf13/viper"
)

// Cli is the walter client to connect to Jira
type Cli struct {
	jira *jira.Client
}

// NewCli returns a command line interface
func NewCli() *Cli {
	base := viper.GetString("JIRA_URL")
	tp := jira.BasicAuthTransport{
		Username: viper.GetString("JIRA_USERNAME"),
		Password: viper.GetString("JIRA_TOKEN"),
	}

	jira, _ := jira.NewClient(tp.Client(), base)

	c := &Cli{
		jira: jira,
	}

	return c
}

// DisplayBoards will render a list of boards
func (c *Cli) DisplayBoards() string {
	boards, err := c.getBoards()
	if err != nil {
		return err.Error()
	}

	ui := ""
	for _, item := range boards {
		ui += fmt.Sprintf("%s\n", item.Name)
	}

	return ui
}

// DisplayIssues will render a list of issues on the board
func (c *Cli) DisplayIssues(boardName, sprintName string) string {
	issues, err := c.getIssues(boardName, sprintName)
	if err != nil {
		return err.Error()
	}

	items := make(map[string][]jira.Issue)

	// Now build a map|slice|array (!) of
	// BoardColumn => Items[]
	for index := 0; index < len(issues); index++ {
		item := issues[index]

		key := item.Fields.Status.Name
		items[key] = append(items[key], item)
	}

	ui := ""
	layout, err := c.getBoardLayout(boardName)
	if err != nil {
		ui += err.Error()
	}

	for _, column := range layout {
		ui += "\n" + column + "\n"
		ui += strings.Repeat("=", len(column)) + "\n"
		for _, v := range items[column] {
			ui += fmt.Sprintf("* %s\n", v.Fields.Summary)
		}
	}
	return ui
}

// DisplayBurndown will render a burndown table for the sprint
func (c *Cli) DisplayBurndown(boardName, sprintName string) string {
	issues, err := c.getIssues(boardName, sprintName)
	if err != nil {
		return err.Error()
	}

	items := make(map[string][]jira.Issue)

	// Now build a map|slice|array (!) of
	// BoardColumn => Isues[]
	for index := 0; index < len(issues); index++ {
		item := issues[index]
		key := item.Fields.Status.Name
		items[key] = append(items[key], item)
	}

	storyFieldKey := fmt.Sprintf("boards.%s.story_point_field", boardName)
	storyField := viper.GetString(storyFieldKey)

	ui := ""

	if storyField == "" {
		ui += fmt.Sprintf("There was no story point field (%s) defined in your configuration file, so cannot calculate points", storyFieldKey)
	}

	layout, err := c.getBoardLayout(boardName)
	if err != nil {
		ui += err.Error()
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 1, 1, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\n", "Column", "Items", "Points")
	fmt.Fprintf(w, "%s\t%s\t%s\n", "------", "-----", "------")
	totalItems := 0
	totalPoints := 0
	for _, column := range layout {
		points := 0
		itemCount := len(items[column])

		for _, item := range items[column] {
			fields, err := c.getIssueCustomFields(item.ID)
			if err != nil {
				points = 0
			}

			v, err := strconv.Atoi(fields[storyField])
			if err != nil {
				v = 0
			}
			points += v
		}
		totalPoints += points
		totalItems += itemCount

		fmt.Fprintf(w, "%s\t%d\t%d\n", column, itemCount, points)
	}
	fmt.Fprintf(w, "%s\t%s\t%s\n", "------", "", "")
	fmt.Fprintf(w, "%s\t%d\t%d\n", "Total", totalItems, totalPoints)
	fmt.Fprintf(w, "%s\t%s\t%s\n", "------", "", "")

	w.Flush() // These two lines need some attention
	return ui // as we flush and return a string
}

// DisplaySprints will render a list of sprints
func (c *Cli) DisplaySprints(boardName string) string {
	sprints, err := c.getSprints(boardName)
	if err != nil {
		return err.Error()
	}

	ui := ""
	for _, sprint := range sprints {
		start := "?"
		end := "?"
		if sprint.StartDate != nil {
			start = sprint.StartDate.Format("02-01-2006")
		}

		if sprint.EndDate != nil {
			end = sprint.EndDate.Format("02-01-2006")
		}
		ui += fmt.Sprintf("* Start: %s End: %s - %s\n", start, end, sprint.Name)
	}

	return ui
}

// getBoards returns all boards
func (c *Cli) getBoards() ([]jira.Board, error) {
	opts := jira.BoardListOptions{}
	list, _, err := c.jira.Board.GetAllBoards(&opts)
	if err != nil {
		return nil, err
	}

	boards := list.Values
	sort.Slice(boards, func(i, j int) bool { return boards[i].Name < boards[j].Name })

	return boards, nil
}

// getBoardLayout will return what the columns are for a given board
func (c *Cli) getBoardLayout(boardName string) ([]string, error) {
	layoutKey := fmt.Sprintf("boards.%s.layout", boardName)
	ok := viper.IsSet(layoutKey)

	if !ok {
		return nil, fmt.Errorf("%s is not defined in the configuration file", layoutKey)
	}

	return viper.GetStringSlice(layoutKey), nil
}

// getBoard returns the board information
func (c *Cli) getBoard(name string) (*jira.Board, error) {
	opts := jira.BoardListOptions{}
	list, _, err := c.jira.Board.GetAllBoards(&opts)
	if err != nil {
		return nil, err
	}

	for _, item := range list.Values {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, fmt.Errorf("unable to find board with name %s", name)
}

// getSprints returns all sprints
func (c *Cli) getSprints(boardName string) ([]jira.Sprint, error) {
	board, err := c.getBoard(boardName)
	if err != nil {
		return nil, err
	}

	sprints, _, err := c.jira.Board.GetAllSprints(strconv.Itoa(board.ID))
	if err != nil {
		return nil, err
	}

	return sprints, nil
}

// getIssues returns all issues
func (c *Cli) getIssues(boardName, sprintName string) ([]jira.Issue, error) {
	board, err := c.getBoard(boardName)
	if err != nil {
		return nil, err
	}

	sprints, _, err := c.jira.Board.GetAllSprints(strconv.Itoa(board.ID))
	if err != nil {
		return nil, err
	}

	sprintID := -1
	for _, sprint := range sprints {
		if sprint.Name == sprintName {
			sprintID = sprint.ID
			break
		}
	}

	issues, _, err := c.jira.Sprint.GetIssuesForSprint(sprintID)
	if err != nil {
		return nil, err
	}

	return issues, nil
}

// getIssueCustomFields returns all custom field data for a given Issue
func (c *Cli) getIssueCustomFields(issueID string) (jira.CustomFields, error) {
	fields, _, err := c.jira.Issue.GetCustomFields(issueID)
	return fields, err
}
