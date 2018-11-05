package cmd

import (
	"fmt"

	"github.com/benmatselby/walter/cli"
	"github.com/spf13/cobra"
)

// boardsListCmd defines the cobra command to list all boards
var boardsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the boards",
	Run: func(cmd *cobra.Command, args []string) {

		c := cli.NewCli()
		boards := c.DisplayBoards()
		fmt.Print(boards)
	},
}

// boardsCmd defines the base "boards" command that allows sub commands
// to hang off
var boardsCmd = &cobra.Command{
	Use:   "boards",
	Short: "Board related commands",
}

// init registers all the commands ultimately to root
func init() {
	boardsCmd.AddCommand(boardsListCmd)
	rootCmd.AddCommand(boardsCmd)
}
