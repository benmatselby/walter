package board

import (
	"fmt"
	"io"
	"os"

	goJira "github.com/andygrunwald/go-jira"
	"github.com/benmatselby/walter/jira"
	"github.com/benmatselby/walter/ui"
	"github.com/spf13/cobra"
)

// NewIssueCommand creates a new `board issues` command
func NewIssueCommand(client jira.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issues",
		Short: "List all the issues for the board",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := ListIssues(client, args[0], os.Stdout)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}

// ListIssues will list all the issues for a given board
func ListIssues(client jira.API, boardName string, w io.Writer) error {
	issues, err := client.GetIssuesForBoard(boardName)
	if err != nil {
		return err
	}

	items := make(map[string][]goJira.Issue)

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
		ui.RenderTitle(w, column)
		for _, v := range items[column] {
			ui.RenderItem(w, v)
		}
	}
	return nil
}
