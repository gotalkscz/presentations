package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// START OMIT
func main() {
	root := &cobra.Command{
		Use: "run",
	}
	root.AddCommand(createCommand())
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "count",
		Short: "počet argumentů",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(len(args))
			return nil
		},
	}
}

// END OMIT
