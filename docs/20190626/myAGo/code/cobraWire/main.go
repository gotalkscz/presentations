package main

import (
	"fmt"
	"os"
	"talk/myAGo/code/service"
	app "talk/myAGo/code/wireFormater"

	"github.com/spf13/cobra"
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
	return &cobra.Command{
		Use: "say",
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := app.CreateApp(service.NewConfig()) // HLnew
			if err != nil {
				return err
			}
			app.Run() // HLnew
			return nil
		},
	}
}

// END OMIT
