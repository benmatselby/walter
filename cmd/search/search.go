package search

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	goJira "github.com/andygrunwald/go-jira"
	"github.com/benmatselby/walter/jira"
	"github.com/benmatselby/walter/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// DefaultMaxResults defines the amount of results we should show as a default
	DefaultMaxResults = 100
)

// CommandOptions defines the options available for searching
type CommandOptions struct {
	Args       []string
	Format     string
	MaxResults int
	Query      string
	Template   string
}

// NewSearchCommand creates a new `search` command
func NewSearchCommand(client jira.API) *cobra.Command {
	var opts CommandOptions

	cmd := &cobra.Command{
		Use:   "search",
		Short: "Search for issues",
		Run: func(cmd *cobra.Command, args []string) {
			opts.Args = args

			err := QueryIssues(client, opts, os.Stdout)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.Format, "format", "f", "list", "The format of the output: list, table")
	flags.IntVar(&opts.MaxResults, "max-results", DefaultMaxResults, "The amount of records to display")
	flags.StringVarP(&opts.Query, "query", "q", "", "The JQL you want to run")
	flags.StringVarP(&opts.Template, "template", "t", "", "The name of the template that as the JQL you want to run")

	return cmd
}

// QueryIssues provides the searching capability
func QueryIssues(client jira.API, opts CommandOptions, w io.Writer) error {
	query := ""
	searchOpts := goJira.SearchOptions{
		MaxResults: opts.MaxResults,
	}

	if opts.Template != "" {
		if !viper.IsSet(fmt.Sprintf("templates.%s", opts.Template)) {
			return fmt.Errorf("%s is not defined", opts.Template)
		}
		count := viper.GetInt(fmt.Sprintf("templates.%s.count", opts.Template))
		query = viper.GetString(fmt.Sprintf("templates.%s.query", opts.Template))
		searchOpts.MaxResults = count
	} else if opts.Query != "" {
		query = opts.Query
	} else {
		return fmt.Errorf("please use --query or --template to search")
	}

	if opts.MaxResults != DefaultMaxResults {
		searchOpts.MaxResults = opts.MaxResults
	}

	issues, err := client.IssueSearch(query, &searchOpts)
	if err != nil {
		return err
	}

	if opts.Format == "table" {
		renderTable(client, issues, w)
	} else {
		renderList(client, issues, w)
	}

	return nil
}

func renderTable(client jira.API, issues []goJira.Issue, w io.Writer) {
	tw := tabwriter.NewWriter(w, 0, 1, 1, ' ', 0)
	fmt.Fprintf(tw, "%s\t%s\n", "Metric", "Count")
	fmt.Fprintf(tw, "%s\t%s\n", "------", "-----")
	totalPoints := 0
	totalUnestimated := 0
	pointsUsed := false

	for _, issue := range issues {
		value, err := jira.GetStoryPoint(issue, "")
		if err != nil {
			totalUnestimated++
		} else {
			totalPoints += value
			pointsUsed = true
		}
	}

	fmt.Fprintf(tw, "%s\t%d\n", "Issues", len(issues))

	if pointsUsed {
		fmt.Fprintf(tw, "%s\t%d\n", "Points", totalPoints)
		fmt.Fprintf(tw, "%s\t%d\n", "Not pointed", totalUnestimated)
	}

	fmt.Fprintf(tw, "%s\t%s\n", "------", "-----")

	tw.Flush()
}

func renderList(client jira.API, issues []goJira.Issue, w io.Writer) {
	for _, issue := range issues {
		ui.RenderItem(w, issue)
	}
}
