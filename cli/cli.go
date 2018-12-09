package cli

import (
	"fmt"
	"sort"
	"strconv"

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
