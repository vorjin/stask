package cmd

import (
	"fmt"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, taskID := range args {
				id, err := strconv.ParseUint(taskID, 10, 64)
				if err != nil {
					return fmt.Errorf("invalid task ID %q: %w", taskID, err)
				}

				task, err := app.DB.DoTask(id)
				if err != nil {
					return fmt.Errorf("failed to complete task %d: %w", id, err)
				}

				color.Green("Task #%d was marked as done!\n", task.ID)
			}
			return nil
		},
	}
}
