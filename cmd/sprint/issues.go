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

// IssueOptions defines what arguments/options the user can provide
type IssueOptions struct {
	Args       []string
	FilterType string
}

// NewIssueCommand creates a new `sprint issues` command
func NewIssueCommand(client jira.API) *cobra.Command {
	var opts IssueOptions

	cmd := &cobra.Command{
		Use:   "issues",
		Short: "List all the issues for the sprint",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			opts.Args = args

			err := ListIssues(client, opts, os.Stdout)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.FilterType, "filter-type", "", "Filter the output based on item type: Story, Sub-task")

	return cmd
}

// ListIssues will list all the issues for a given board and sprint
func ListIssues(client jira.API, opts IssueOptions, w io.Writer) error {
	boardName := opts.Args[0]
	sprintName := opts.Args[1]
	issues, err := client.GetIssues(boardName, sprintName)
	if err != nil {
		return err
	}

	items := make(map[string][]goJira.Issue)

	// Now build a map|slice|array (!) of
	// BoardColumn => Items[]
	for index := 0; index < len(issues); index++ {
		item := issues[index]

		if opts.FilterType != "" && opts.FilterType != item.Fields.Type.Name {
			continue
		}

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
			fmt.Fprintf(w, "* %s - %s\n", v.Key, v.Fields.Summary)
		}
	}
	return nil
}
