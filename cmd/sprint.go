package cmd

import (
	"fmt"
	"os"
	"strconv"

	jira "github.com/andygrunwald/go-jira"
	"github.com/spf13/cobra"
)

var sprintIssuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "List all the issues for the sprint",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		boardName := args[0]
		sprintName := args[1]

		issues, err := getIssues(boardName, sprintName)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(2)
		}

		for _, issue := range issues {
			fmt.Println(fmt.Sprintf("* %s", issue.Fields.Summary))
		}
	},
}

var sprintListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the sprints",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		boardName := args[0]

		board, err := getBoard(boardName)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(2)
		}

		client := getClient()
		sprints, _, err := client.Board.GetAllSprints(strconv.Itoa(board.ID))
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(2)
		}

		for _, sprint := range sprints {
			start := "?"
			end := "?"
			if sprint.StartDate != nil {
				start = sprint.StartDate.Format("02-01-2006")
			}

			if sprint.EndDate != nil {
				end = sprint.EndDate.Format("02-01-2006")
			}
			fmt.Println(fmt.Sprintf("* Start: %s End: %s - %s", start, end, sprint.Name))
		}
	},
}

var sprintCmd = &cobra.Command{
	Use:   "sprint",
	Short: "Sprint related commands",
}

func init() {
	sprintCmd.AddCommand(sprintIssuesCmd)
	sprintCmd.AddCommand(sprintListCmd)
	rootCmd.AddCommand(sprintCmd)
}

func getIssues(boardName, sprintName string) ([]jira.Issue, error) {
	client := getClient()

	board, err := getBoard(boardName)
	if err != nil {
		return nil, err
	}

	sprints, _, err := client.Board.GetAllSprints(strconv.Itoa(board.ID))
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

	issues, _, err := client.Sprint.GetIssuesForSprint(sprintID)
	if err != nil {
		return nil, err
	}

	return issues, nil
}
