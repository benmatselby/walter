package sprint

import (
	"fmt"

	"github.com/benmatselby/walter/cli"
	"github.com/spf13/cobra"
)

// NewSprintIssueCommand creates a new `sprint issues` command
func NewSprintIssueCommand(cli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issues",
		Short: "List all the issues for the sprint",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			boardName := args[0]
			sprintName := args[1]
			issues := cli.DisplayIssues(boardName, sprintName)
			fmt.Print(issues)
		},
	}

	return cmd
}
