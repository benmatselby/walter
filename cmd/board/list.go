package board

import (
	"fmt"
	"io"
	"os"

	"github.com/benmatselby/walter/jira"
	"github.com/spf13/cobra"
)

// NewListCommand creates a new `board list` command
func NewListCommand(client jira.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all the boards",
		Run: func(cmd *cobra.Command, args []string) {
			err := DisplayBoards(client, os.Stdout)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}

// DisplayBoards will render the boards you have access to
func DisplayBoards(client jira.API, w io.Writer) error {
	boards, err := client.GetBoards()
	if err != nil {
		return err
	}

	for _, item := range boards {
		fmt.Fprintf(w, "%s\n", item.Name)
	}

	return nil
}
