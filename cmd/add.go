package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"stask/db"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task to your list",
	Long:  `Add a new task to your list. The task description can be a single word or a sentence.`,
	Run: func(cmd *cobra.Command, args []string) {
		todoTask := []byte(strings.Join(args, " "))

		err := db.AddToDOTask(todoTask)

		if err != nil {
			panic(err)
		}

		color.Green("Task was added succesfully.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
