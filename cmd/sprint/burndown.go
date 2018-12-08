package sprint

import (
	"fmt"

	"github.com/benmatselby/walter/cli"
	"github.com/spf13/cobra"
)

// NewSprintBurndownCommand creates a new `sprint burndown` command
func NewSprintBurndownCommand(cli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burndown",
		Short: "Display the burndown for the sprint",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			boardName := args[0]
			sprintName := args[1]
			burndown := cli.DisplayBurndown(boardName, sprintName)
			fmt.Print(burndown)
		},
	}

	return cmd
}
