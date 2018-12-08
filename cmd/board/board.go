package board

import (
	"github.com/benmatselby/walter/cli"
	"github.com/benmatselby/walter/jira"
	"github.com/spf13/cobra"
)

// NewBoardCommand creates a new `board` command
func NewBoardCommand(cli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "board",
		Short: "Board related commands",
	}

	// Temporary - This will be a param when all converted over
	client := jira.NewClient()
	// Temporary

	cmd.AddCommand(
		NewListCommand(&client),
	)
	return cmd
}
