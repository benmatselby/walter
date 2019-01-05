package jira

import (
	"fmt"
	"sort"
	"strconv"

	jira "github.com/andygrunwald/go-jira"
	"github.com/spf13/viper"
)

// API defines the interface
type API interface {
	GetBoards() ([]jira.Board, error)
	GetBoard(name string) (*jira.Board, error)
	GetBoardLayout(name string) ([]string, error)
	GetSprints(boardName string) ([]jira.Sprint, error)
	GetIssues(boardName, sprintName string) ([]jira.Issue, error)
	GetIssuesForBoard(boardName string) ([]jira.Issue, error)
	GetIssueCustomFields(issueID string) (jira.CustomFields, error)
}

// Client is the concrete implementation of the API interface
type Client struct {
	jira *jira.Client
}

// NewClient will return a new internal jira client. This only defines what the application uses
func NewClient() Client {
	base := viper.GetString("JIRA_URL")
	tp := jira.BasicAuthTransport{
		Username: viper.GetString("JIRA_USERNAME"),
		Password: viper.GetString("JIRA_TOKEN"),
	}

	jira, _ := jira.NewClient(tp.Client(), base)

	client := Client{
		jira: jira,
	}

	return client
}

// GetBoards will return the boards you can access
func (c *Client) GetBoards() ([]jira.Board, error) {
	opts := jira.BoardListOptions{}
	list, _, err := c.jira.Board.GetAllBoards(&opts)
	if err != nil {
		return nil, err
	}

	boards := list.Values
	sort.Slice(boards, func(i, j int) bool { return boards[i].Name < boards[j].Name })

	return boards, nil
}

// GetBoard will return a single board given a name
func (c *Client) GetBoard(name string) (*jira.Board, error) {
	list, err := c.GetBoards()
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, fmt.Errorf("unable to find board with name %s", name)
}

// GetBoardLayout will return what the columns are for a given board
func (c *Client) GetBoardLayout(name string) ([]string, error) {
	layoutKey := fmt.Sprintf("boards.%s.layout", name)
	ok := viper.IsSet(layoutKey)

	if !ok {
		return nil, fmt.Errorf("%s is not defined in the configuration file", layoutKey)
	}

	return viper.GetStringSlice(layoutKey), nil
}

// GetSprints will return a list of sprints
func (c *Client) GetSprints(boardName string) ([]jira.Sprint, error) {
	board, err := c.GetBoard(boardName)
	if err != nil {
		return nil, err
	}

	sprints, _, err := c.jira.Board.GetAllSprints(strconv.Itoa(board.ID))
	if err != nil {
		return nil, err
	}

	return sprints, nil
}

// GetIssues will return a list of issues for a given board and sprint
func (c *Client) GetIssues(boardName, sprintName string) ([]jira.Issue, error) {
	board, err := c.GetBoard(boardName)
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

// searchResult is pulled from the jira library
// It's not a public struct, but needed the struct for the Do function in GetIssuesForBoard
type searchResult struct {
	Issues     []jira.Issue `json:"issues" structs:"issues"`
	StartAt    int          `json:"startAt" structs:"startAt"`
	MaxResults int          `json:"maxResults" structs:"maxResults"`
	Total      int          `json:"total" structs:"total"`
}

// GetIssuesForBoard will return a list of issues for a given board
func (c *Client) GetIssuesForBoard(boardName string) ([]jira.Issue, error) {
	board, err := c.GetBoard(boardName)
	if err != nil {
		return nil, err
	}

	// I cannot find a method in the go-jira package to do this, but I also struggled
	// to find the API in the Jira documentation. Found this in
	// https://community.atlassian.com/t5/Answers-Developer-Questions/Retrieve-all-issues-from-a-kanban-board-using-JIRA-rest-api/qaq-p/538719
	req, _ := c.jira.NewRequest("GET", fmt.Sprintf("rest/agile/latest/board/%v/issue", board.ID), nil)

	result := new(searchResult)
	_, err = c.jira.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result.Issues, nil
}

// GetIssueCustomFields returns all custom field data for a given Issue
func (c *Client) GetIssueCustomFields(issueID string) (jira.CustomFields, error) {
	fields, _, err := c.jira.Issue.GetCustomFields(issueID)
	return fields, err
}
