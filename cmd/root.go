package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/benmatselby/walter/cli"
	"github.com/benmatselby/walter/cmd/sprint"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once
func Execute() {
	// We need the configuration loaded before we create a NewCli
	// as that needs the viper configration up and running
	initConfig()

	// Build a new Cli
	cli := cli.NewCli()

	// Build the root command
	cmd := NewRootCommand(cli)

	// Execute the application
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// NewRootCommand builds the main cli application and
// adds all the child commands
func NewRootCommand(cli *cli.Cli) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "walter",
		Short: "CLI application for retrieving data from Jira",
	}

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.walter/config.yaml)")

	cmd.AddCommand(
		NewBoardCommand(cli),
		sprint.NewSprintCommand(cli),
	)

	return cmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".walter" (without extension).
		path := strings.Join([]string{home, ".walter"}, "/")
		viper.AddConfigPath(path)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config: %s\n", err)
	}
}
