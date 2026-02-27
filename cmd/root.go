// Package cmd is for managing cobra commands
package cmd

import (
	"os"

	"stask/db"

	"github.com/spf13/cobra"
)

type App struct {
	DB db.TaskStore
}

func Execute(database db.TaskStore) {
	app := &App{DB: database}

	rootCmd := &cobra.Command{
		Use:   "task",
		Short: "task is a CLI for managing your TODOs.",
		Long: `task is a CLI for managing your TODOs.
		It allows you to add, list, and mark tasks as complete.`,
	}

	rootCmd.AddCommand(app.NewAddCmd())
	rootCmd.AddCommand(app.NewCompletedCmd())
	rootCmd.AddCommand(app.NewDoCmd())
	rootCmd.AddCommand(app.NewListCmd())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
