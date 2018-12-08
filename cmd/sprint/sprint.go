package sprint

import (
	"github.com/benmatselby/walter/cli"
	"github.com/spf13/cobra"
)

// NewSprintCommand creates a new `sprint` command
func NewSprintCommand(cli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sprint",
		Short: "Sprint related commands",
	}

	cmd.AddCommand(
		NewSprintListCommand(cli),
		NewSprintBurndownCommand(cli),
		NewSprintIssueCommand(cli),
	)

	return cmd
}
