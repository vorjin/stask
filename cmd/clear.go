package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func (app *App) NewClearCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "clear",
		Short: "clears TODO tasks list",
		Run: func(cmd *cobra.Command, args []string) {
			err := app.DB.DeleteTasksBucket()
			if err != nil {
				fmt.Printf("Error clearing TODO list. Err: %v", err)
				os.Exit(1)
			}

			color.Green("TODO list was succesfully cleared!")
		},
	}
}
