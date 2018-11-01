package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/viper"

	jira "github.com/andygrunwald/go-jira"
	"github.com/spf13/cobra"
)

// boardsCmd represents the boards command
var boardsCmd = &cobra.Command{
	Use:   "boards",
	Short: "List all the boards in Jira",
	Run: func(cmd *cobra.Command, args []string) {
		base := viper.GetString("JIRA_URL")
		tp := jira.BasicAuthTransport{
			Username: viper.GetString("JIRA_USERNAME"),
			Password: viper.GetString("JIRA_TOKEN"),
		}

		client, err := jira.NewClient(tp.Client(), base)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(2)
		}

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

func init() {
	rootCmd.AddCommand(boardsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// boardsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// boardsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
