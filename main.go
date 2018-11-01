package main

import (
	"fmt"
	"os"

	jira "github.com/andygrunwald/go-jira"
)

func main() {
	base := os.Getenv("JIRA_URL")
	tp := jira.BasicAuthTransport{
		Username: os.Getenv("JIRA_USERNAME"),
		Password: os.Getenv("JIRA_TOKEN"),
	}

	client, err := jira.NewClient(tp.Client(), base)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}

	opts := jira.BoardListOptions{}
	list, _, err := client.Board.GetAllBoards(&opts)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}

	for _, item := range list.Values {
		fmt.Println(item.Name)
	}
}
