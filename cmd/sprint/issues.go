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
	MaxResults int
}

// NewIssueCommand creates a new `sprint issues` command
func NewIssueCommand(client jira.API) *cobra.Command {
	var opts IssueOptions

	cmd := &cobra.Command{
		Use:   "issues",
		Short: "List all the issues for a given project sprint",
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
	flags.IntVar(&opts.MaxResults, "max-results", 100, "The amount of records to display")

	return cmd
}

// ListIssues will list all the issues for a given board and sprint
func ListIssues(client jira.API, opts IssueOptions, w io.Writer) error {
	boardName := opts.Args[0]
	sprintName := opts.Args[1]

	query := fmt.Sprintf("sprint = '%s'", sprintName)

	if len(opts.Args) > 2 {
		projectName := opts.Args[2]
		query += fmt.Sprintf(" and project = '%s'", projectName)
	}
	
	if opts.FilterType != "" {
		query += fmt.Sprintf(" and type = '%s'", opts.FilterType)
	}

	searchOpts := goJira.SearchOptions{
		MaxResults: opts.MaxResults,
	}
	issues, err := client.IssueSearch(query, &searchOpts)

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
			fmt.Fprintf(w, "* %s - %s\n", v.Key, v.Fields.Summary)
		}
	}
	return nil
}
