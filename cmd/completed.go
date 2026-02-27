package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func (app *App) NewCompletedCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completed",
		Short: "show completed tasks",
		Long:  "show completed tasks",
		Run: func(cmd *cobra.Command, args []string) {
			hours, err := cmd.Flags().GetInt("time")
			if err != nil {
				fmt.Printf("Error parsing time flag. Err: %v", err)
				os.Exit(1)
			}

			tasks, err := app.DB.ListCompletedTasks(hours)
			if err != nil {
				fmt.Printf("Error getting completed tasks. Err: %v", err)
				os.Exit(1)
			}

			if len(tasks) == 0 {
				color.Red("There are no completed tasks! Maaaan, time to do something..\n")
				return
			}

			color.Cyan("This are your completed tasks for the last %d hours: \n", hours)
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task.Name)
			}
		},
	}
	cmd.Flags().IntP("time", "t", 24, "how many hours ago tasks were completed")
	return cmd
}
