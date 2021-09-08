package cmd

import (
	"github.com/spf13/cobra"
)

var (
	searchCommand = &cobra.Command{
		Use:   "search",
		Short: "Search the path of the specified file",
		Long:  "Search the path of the specified file",
		RunE:  run,
	}
	strategy string
)

func init() {
	rootCmd.AddCommand(searchCommand)
	searchCommand.Flags().StringVarP(&strategy, "strategy", "s", "", "choose the algorithm for searching")
}

func run(cmd *cobra.Command, args []string) error {
	return nil
}
