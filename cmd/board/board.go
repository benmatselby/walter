package board

import (
	"github.com/benmatselby/walter/jira"
	"github.com/spf13/cobra"
)

// NewBoardCommand creates a new `board` command
func NewBoardCommand(client jira.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "board",
		Short: "Board related commands",
	}

	cmd.AddCommand(
		NewListCommand(client),
	)
	return cmd
}
