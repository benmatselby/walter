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
	GetSprints(boardName string) ([]jira.Sprint, error)
}

// Client is the concrete implemntation of the API interface
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
