package sprint

import (
	"fmt"
	"io"
	"os"

	"github.com/benmatselby/walter/jira"
	"github.com/spf13/cobra"
)

// NewListCommand creates a new `sprint list` command
func NewListCommand(client jira.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List all the sprints",
		Example: "walter sprint list \"my board\"",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := ListSprints(client, args[0], os.Stdout)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}

// ListSprints will list all the sprints for a given board
func ListSprints(client jira.API, boardName string, w io.Writer) error {
	sprints, err := client.GetSprints(boardName)
	if err != nil {
		return err
	}

	for _, sprint := range sprints {
		fmt.Fprintf(w, "%s\n", sprint.Name)
	}

	return nil
}
