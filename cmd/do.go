package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func (app *App) NewDoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "do [task number]",
		Short: "Mark a task as complete",
		Long:  `Mark a task as complete by providing its number from the list.`,
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := app.DB.DoTask(args)
			if err != nil {
				fmt.Printf("Error marking task(s) as 'Done'. Err: %v", err)
				os.Exit(1)
			}

			for _, task := range tasks {
				color.Green("Task #%s was marked as done!\n", task.Name)
			}
		},
	}
}
