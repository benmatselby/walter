package sprint

import (
	"fmt"
	"io"
	"os"

	goJira "github.com/andygrunwald/go-jira"
	"github.com/benmatselby/walter/jira"
	"github.com/benmatselby/walter/ui"
	"github.com/spf13/cobra"
)

// IssueOptions defines what arguments/options the user can provide
type IssueOptions struct {
	Args       []string
	FilterType string
	GroupBy    string
	MaxResults int
}

const (
	// GroupByBoard allows users to group the output by the Jira board.
	GroupByBoard = "board"

	// GroupByLabel allows the users to group items by the labels assigned to
	// the item.
	GroupByLabel = "label"
)

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
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.FilterType, "filter-type", "", "Filter the output based on item type: Story, Sub-task")
	flags.StringVar(&opts.GroupBy, "group-by", "board", fmt.Sprintf("Group the work items: %s, %s", GroupByBoard, GroupByLabel))
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

	if opts.GroupBy == GroupByBoard {
		return renderByBoard(client, w, boardName, issues)
	} else if opts.GroupBy == GroupByLabel {
		return renderByTag(client, w, boardName, issues)
	} else {
		return fmt.Errorf("%s is not a valid grouping", opts.GroupBy)
	}
}

func renderByBoard(client jira.API, w io.Writer, boardName string, issues []goJira.Issue) error {
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
		ui.RenderTitle(w, column)
		for _, item := range items[column] {
			ui.RenderItem(w, item)
		}
	}

	return nil
}

func renderByTag(client jira.API, w io.Writer, boardName string, issues []goJira.Issue) error {
	labels := make(map[string][]goJira.Issue)

	for _, item := range issues {
		for _, label := range item.Fields.Labels {
			labels[label] = append(labels[label], item)
		}
	}

	for label, items := range labels {
		ui.RenderTitle(w, label)

		for _, item := range items {
			ui.RenderItem(w, item)
		}
	}

	return nil
}
