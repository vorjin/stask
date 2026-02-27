package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func (app *App) NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all of your todo tasks",
		Long:  `List all of your todo tasks currently stored in the database.`,
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := app.DB.ListToDoTasks()
			if err != nil {
				fmt.Printf("Error listing TODO tasks. Err: %v", err)
				os.Exit(1)
			}

			if len(tasks) == 0 {
				color.Green("There are currently no tasks to do. Go for a walk :D\n")
				return
			}

			color.Magenta("This are your tasks: \n")

			for _, task := range tasks {
				fmt.Printf("%d. %s\n", task.ID, task.Name)
			}
		},
	}
}
