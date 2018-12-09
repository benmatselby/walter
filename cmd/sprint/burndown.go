package sprint

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"text/tabwriter"

	goJira "github.com/andygrunwald/go-jira"
	"github.com/benmatselby/walter/jira"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewBurndownCommand creates a new `sprint burndown` command
func NewBurndownCommand(client jira.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burndown",
		Short: "Display the burndown for the sprint",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			err := ShowBurndown(client, args[0], args[1], os.Stdout)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}

// ShowBurndown will provide burndown data in a tabular format
func ShowBurndown(client jira.API, boardName, sprintName string, w io.Writer) error {
	issues, err := client.GetIssues(boardName, sprintName)
	if err != nil {
		return err
	}

	items := make(map[string][]goJira.Issue)

	// Now build a map|slice|array (!) of
	// BoardColumn => Isues[]
	for index := 0; index < len(issues); index++ {
		item := issues[index]
		key := item.Fields.Status.Name
		items[key] = append(items[key], item)
	}

	storyFieldKey := fmt.Sprintf("boards.%s.story_point_field", boardName)
	storyField := viper.GetString(storyFieldKey)

	ui := ""

	if storyField == "" {
		ui += fmt.Sprintf("There was no story point field (%s) defined in your configuration file, so cannot calculate points", storyFieldKey)
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
			fields, err := client.GetIssueCustomFields(item.ID)
			if err != nil {
				points = 0
			}

			v, err := strconv.Atoi(fields[storyField])
			if err != nil {
				v = 0
			}
			points += v
		}
		totalPoints += points
		totalItems += itemCount

		fmt.Fprintf(tw, "%s\t%d\t%d\n", column, itemCount, points)
	}
	fmt.Fprintf(tw, "%s\t%s\t%s\n", "------", "", "")
	fmt.Fprintf(tw, "%s\t%d\t%d\n", "Total", totalItems, totalPoints)
	fmt.Fprintf(tw, "%s\t%s\t%s\n", "------", "", "")

	tw.Flush()

	return nil
}
