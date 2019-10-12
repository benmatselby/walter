package sprint

import (
	"fmt"
	"io"
	"os"

	"github.com/benmatselby/walter/jira"
	"github.com/spf13/cobra"
)

// ListSprintOptions defines the options available for listing the sprints
type ListSprintOptions struct {
	Args    []string
	Verbose bool
}

// NewListCommand creates a new `sprint list` command
func NewListCommand(client jira.API) *cobra.Command {
	var opts ListSprintOptions

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List all the sprints",
		Example: "walter sprint list \"my board\"",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			opts.Args = args

			err := ListSprints(client, opts, os.Stdout)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.Verbose, "verbose", "v", false, "Verbose output")

	return cmd
}

// ListSprints will list all the sprints for a given board
func ListSprints(client jira.API, opts ListSprintOptions, w io.Writer) error {
	sprints, err := client.GetSprints(opts.Args[0])
	if err != nil {
		return err
	}

	for _, sprint := range sprints {
		fmt.Fprintf(w, "%s", sprint.Name)
		if opts.Verbose {
			fmt.Fprintln(w)
			fmt.Fprintf(w, " ID: %v\n", sprint.ID)
			fmt.Fprintf(w, " Start date: %v\n", sprint.StartDate)
			fmt.Fprintf(w, " End date: %v\n", sprint.EndDate)
			fmt.Fprintf(w, " Completed date: %v\n", sprint.CompleteDate)
		}
		fmt.Fprintln(w)
	}

	return nil
}
