package sprint

import (
	"fmt"
	"io"
	"os"
	"strings"

	goJira "github.com/andygrunwald/go-jira"
	"github.com/benmatselby/walter/jira"
	"github.com/spf13/cobra"
)

// NewIssueCommand creates a new `sprint issues` command
func NewIssueCommand(client jira.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issues",
		Short: "List all the issues for the sprint",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			err := ListIssues(client, args[0], args[1], os.Stdout)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}

// ListIssues will list all the issues for a given board and sprint
func ListIssues(client jira.API, boardName, sprintName string, w io.Writer) error {
	issues, err := client.GetIssues(boardName, sprintName)
	if err != nil {
		return err
	}

	items := make(map[string][]goJira.Issue)

	// Now build a map|slice|array (!) of
	// BoardColumn => Items[]
	for index := 0; index < len(issues); index++ {
		item := issues[index]

		key := item.Fields.Status.Name
		items[key] = append(items[key], item)
	}

	layout, err := client.GetBoardLayout(boardName)
	if err != nil {
		return err
	}

	for _, column := range layout {
		fmt.Fprintf(w, "\n%s\n", column)
		fmt.Fprintf(w, strings.Repeat("=", len(column))+"\n")
		for _, v := range items[column] {
			fmt.Fprintf(w, "* %s\n", v.Fields.Summary)
		}
	}
	return nil
}
