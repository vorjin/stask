package cmd

import (
	"fmt"
	"os"
	"strings"

	"stask/db"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task to your list",
	Long:  `Add a new task to your list. The task description can be a single word or a sentence.`,
	Run: func(cmd *cobra.Command, args []string) {
		todoTask := []byte(strings.Join(args, " "))

		err := db.AddToDoTask(todoTask)
		if err != nil {
			fmt.Printf("Error adding new task. Err: %v", err)
			os.Exit(1)
		}

		color.Green("Task was added succesfully.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
