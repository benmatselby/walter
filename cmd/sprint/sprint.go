package sprint

import (
	"github.com/benmatselby/walter/cli"
	"github.com/benmatselby/walter/jira"
	"github.com/spf13/cobra"
)

// NewSprintCommand creates a new `sprint` command
func NewSprintCommand(cli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sprint",
		Short: "Sprint related commands",
	}

	// Temporary - This will be a param when all converted over
	client := jira.NewClient()
	// Temporary

	cmd.AddCommand(
		NewListCommand(&client),
		NewSprintBurndownCommand(cli),
		NewSprintIssueCommand(cli),
	)

	return cmd
}
