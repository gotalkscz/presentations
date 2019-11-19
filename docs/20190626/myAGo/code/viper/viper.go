package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	root := &cobra.Command{}
	root.AddCommand(createCommand())
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// START OMIT
func createCommand() *cobra.Command {
	v := viper.New()
	v.SetDefault("format", "%d")
	v.AutomaticEnv()
	cmd := &cobra.Command{
		Use:   "count",
		Short: "počet argumentů",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf(v.GetString("format"), len(args))
			return nil
		},
	}
	return cmd
}

// END OMIT
