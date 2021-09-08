package cmd

import (
	pkgFinder "file-finder/pkg/lib/finder"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	searchCommand = &cobra.Command{
		Use:   "search",
		Short: "Search the path of the specified file",
		Long:  "Search the path of the specified file",
		RunE:  run,
	}
	strategy  string
	directory string
	trace     bool
)

func init() {
	rootCmd.AddCommand(searchCommand)
	searchCommand.Flags().StringVarP(&strategy, "strategy", "s", "", "choose the algorithm for searching")
	searchCommand.Flags().StringVarP(&directory, "directory", "d", "", "specify the directory")
	searchCommand.Flags().BoolVarP(&trace, "trace-log", "", false, "Viewing the trace logs")
}

func run(cmd *cobra.Command, args []string) error {
	var fileName string
	fmt.Printf("\nPlease enter the filename: ")
	fmt.Scanf("%s", &fileName)
	if fileName == "" {
		return fmt.Errorf("invalid filename")
	}

	if trace {
		fmt.Printf("===trace logs=== \n\n")
	}

	finder := pkgFinder.NewFinder(directory).ViewTrace(trace)
	path, err := finder.FindPath(fileName)
	if err != nil {
		return fmt.Errorf("failed to find the file: %v", err)
	}

	if path == "" {
		fmt.Printf("\nfile not found \n\n")
		return nil
	}

	fmt.Printf("\n===search result=== \n\n")
	fmt.Printf("Location: %s \n\n", path)
	return nil
}
