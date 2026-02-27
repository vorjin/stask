package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func (app *App) NewDoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "do [task number]",
		Short: "Mark a task as complete",
		Long:  `Mark a task as complete by providing its number from the list.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, taskID := range args {
				id, err := strconv.ParseUint(taskID, 10, 64)
				if err != nil {
					fmt.Printf("Error parsing args into Task IDs. Err: %v", err)
					os.Exit(1)
				}

				task, err := app.DB.DoTask(id)
				if err != nil {
					fmt.Printf("Error marking task(s) as 'Done'. Err: %v", err)
					os.Exit(1)
				}

				color.Green("Task #%d was marked as done!\n", task.ID)
			}
		},
	}
}
