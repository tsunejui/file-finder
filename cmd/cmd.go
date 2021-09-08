package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "finder",
		Short: "A CLI tool for finding the file",
		Long:  "This tool provides an easy and extensible way to finding the file.",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}
