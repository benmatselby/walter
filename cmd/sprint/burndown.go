package sprint

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	goJira "github.com/andygrunwald/go-jira"
	"github.com/benmatselby/walter/jira"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// BurndownOptions defines what arguments/options the user can provide
type BurndownOptions struct {
	Args       []string
	FilterType string
	MaxResults int
}

// NewBurndownCommand creates a new `sprint burndown` command
func NewBurndownCommand(client jira.API) *cobra.Command {
	var opts BurndownOptions

	cmd := &cobra.Command{
		Use:   "burndown",
		Short: "Display the burndown for the sprint",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			opts.Args = args
			err := ShowBurndown(client, opts, os.Stdout)
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

// ShowBurndown will provide burndown data in a tabular format
func ShowBurndown(client jira.API, opts BurndownOptions, w io.Writer) error {
	boardName := opts.Args[0]
	sprintName := opts.Args[1]

	query := fmt.Sprintf("sprint = '%s'", sprintName)

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

	for index := 0; index < len(issues); index++ {
		item := issues[index]
		key := item.Fields.Status.Name
		items[key] = append(items[key], item)
	}

	storyField := viper.GetString(fmt.Sprintf("boards.%s.story_point_field", boardName))
	storyFields := viper.GetStringSlice(fmt.Sprintf("boards.%s.story_point_fields", boardName))

	ui := ""

	if storyField == "" && len(storyFields) == 0 {
		ui += fmt.Sprintf("There was no story point field(s) defined in your configuration file, so cannot calculate points")
	}

	layout, err := client.GetBoardLayout(boardName)
	if err != nil {
		ui += err.Error()
	}

	tw := tabwriter.NewWriter(w, 0, 1, 1, ' ', 0)
	fmt.Fprintf(tw, "%s\t%s\t%s\n", "Column", "Items", "Points")
	fmt.Fprintf(tw, "%s\t%s\t%s\n", "------", "-----", "------")
	totalItems := 0
	totalPoints := 0
	for _, column := range layout {
		points := 0
		itemCount := len(items[column])

		for _, item := range items[column] {
			value := item.Fields.Unknowns[storyField]
			if value != nil {
				points += int(value.(float64))
			} else {
				for _, key := range storyFields {
					value := item.Fields.Unknowns[key]
					if value != nil {
						points += int(value.(float64))
						break
					}
				}
			}
		}
		totalPoints += points
		totalItems += itemCount

		fmt.Fprintf(tw, "%s\t%d\t%d\n", column, itemCount, points)
	}
	fmt.Fprintf(tw, "%s\t%s\t%s\n", "------", "-----", "------")
	fmt.Fprintf(tw, "%s\t%d\t%d\n", "Total", totalItems, totalPoints)
	fmt.Fprintf(tw, "%s\t%s\t%s\n", "------", "-----", "------")

	tw.Flush()

	return nil
}
