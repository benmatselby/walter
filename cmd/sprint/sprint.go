package sprint

import (
	"github.com/benmatselby/walter/jira"
	"github.com/spf13/cobra"
)

// NewSprintCommand creates a new `sprint` command
func NewSprintCommand(client jira.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sprint",
		Short: "Sprint related commands",
	}

	cmd.AddCommand(
		NewListCommand(client),
		NewBurndownCommand(client),
		NewIssueCommand(client),
	)

	return cmd
}
