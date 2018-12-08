package sprint

import (
	"fmt"

	"github.com/benmatselby/walter/cli"
	"github.com/spf13/cobra"
)

// NewSprintListCommand creates a new `sprint list` command
func NewSprintListCommand(cli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List all the sprints",
		Example: "walter sprint list \"my board\"",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			boardName := args[0]
			sprints := cli.DisplaySprints(boardName)
			fmt.Print(sprints)
		},
	}

	return cmd
}
