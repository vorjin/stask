package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func (app *App) NewAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add [task description]",
		Short: "Add a new task to your list",
		Long:  `Add a new task to your list. The task description can be a single word or a sentence.`,
		Run: func(cmd *cobra.Command, args []string) {
			taskName := strings.Join(args, " ")

			err := app.DB.AddTask(taskName)
			if err != nil {
				fmt.Printf("Error adding new task. Err: %v\n", err)
				os.Exit(1)
			}

			color.Green("Task was added successfully.")
		},
	}
}
