package cmd

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

// NewSprintBurndownCommand creates a new `sprint burndown` command
func NewSprintBurndownCommand(cli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burndown",
		Short: "Display the burndown for the sprint",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			boardName := args[0]
			sprintName := args[1]
			burndown := cli.DisplayBurndown(boardName, sprintName)
			fmt.Print(burndown)
		},
	}

	return cmd
}

// NewSprintListCommand creates a new `sprint list` command
func NewSprintListCommand(cli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all the sprints",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			boardName := args[0]
			sprints := cli.DisplaySprints(boardName)
			fmt.Print(sprints)
		},
	}

	return cmd
}

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
