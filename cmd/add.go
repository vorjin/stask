package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func (app *App) NewAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add [task description]",
		Short: "Add a new task to your list",
		Long:  `Add a new task to your list. The task description can be a single word or a sentence.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			taskName := strings.Join(args, " ")

			err := app.DB.AddTask(taskName)
			if err != nil {
				return fmt.Errorf("failed to add task: %w", err)
			}

			color.Green("Task was added successfully.")
			return nil
		},
	}
}
