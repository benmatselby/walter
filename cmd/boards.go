package cmd

import (
	"fmt"

	"github.com/benmatselby/walter/cli"
	"github.com/spf13/cobra"
)

// NewBoardListCommand creates a new `board list` command
func NewBoardListCommand(cli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all the boards",
		Run: func(cmd *cobra.Command, args []string) {
			boards := cli.DisplayBoards()
			fmt.Println(boards)
		},
	}

	return cmd
}

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
