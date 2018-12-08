package board

import (
	"github.com/benmatselby/walter/cli"
	"github.com/spf13/cobra"
)

// NewBoardCommand creates a new `board` command
func NewBoardCommand(cli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "board",
		Short: "Board related commands",
	}
	cmd.AddCommand(
		NewBoardListCommand(cli),
	)
	return cmd
}
