package cmd

import (
	"fmt"
	"os"
	"sort"

	jira "github.com/andygrunwald/go-jira"
	"github.com/spf13/cobra"
)

var boardsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the boards",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClient()

		opts := jira.BoardListOptions{}
		list, _, err := client.Board.GetAllBoards(&opts)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(2)
		}

		boards := list.Values
		sort.Slice(boards, func(i, j int) bool { return boards[i].Name < boards[j].Name })

		for _, item := range boards {
			fmt.Println(item.Name)
		}
	},
}

var boardsCmd = &cobra.Command{
	Use:   "boards",
	Short: "Board related commands",
}

func init() {
	boardsCmd.AddCommand(boardsListCmd)
	rootCmd.AddCommand(boardsCmd)
}

func getBoard(name string) (*jira.Board, error) {
	client := getClient()
	opts := jira.BoardListOptions{}
	list, _, err := client.Board.GetAllBoards(&opts)
	if err != nil {
		return nil, err
	}

	for _, item := range list.Values {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, fmt.Errorf("unable to find board with name %s", name)
}
