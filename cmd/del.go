package cmd

import (
	"fmt"
	"strconv"

	"stask/model"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func (app *App) NewDelCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "del [task number]",
		Short: "Delete a task from the list",
		Long:  `Delete a task from the list by providing its number from the list.`,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, taskID := range args {
				id, err := strconv.ParseUint(taskID, 10, 64)
				if err != nil {
					return fmt.Errorf("invalid task ID %q: %w", taskID, err)
				}

				task, err := app.DB.UpdateTask(model.Deleted, id)
				if err != nil {
					return fmt.Errorf("failed to delete task %d: %w", id, err)
				}

				color.Red("Task #%d was deleted.\n", task.ID)
			}
			return nil
		},
	}
}
