package cmd

import (
	"fmt"

	"github.com/benmatselby/walter/cli"
	"github.com/spf13/cobra"
)

// sprintIssuesCmd defines the cobra command to list all issues
var sprintIssuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "List all the issues for the sprint",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		boardName := args[0]
		sprintName := args[1]

		c := cli.NewCli()
		issues := c.DisplayIssues(boardName, sprintName)
		fmt.Print(issues)
	},
}

// sprintListCmd defines the cobra command to list all the sprints
var sprintListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the sprints",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		boardName := args[0]

		c := cli.NewCli()
		sprints := c.DisplaySprints(boardName)
		fmt.Print(sprints)
	},
}

// sprintCmd defines the base "sprint" command that allows sub commands
// to hang off
var sprintCmd = &cobra.Command{
	Use:   "sprint",
	Short: "Sprint related commands",
}

// init registers all the commands ultimately to root
func init() {
	sprintCmd.AddCommand(sprintIssuesCmd)
	sprintCmd.AddCommand(sprintListCmd)
	rootCmd.AddCommand(sprintCmd)
}
