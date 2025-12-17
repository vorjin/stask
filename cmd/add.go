package cmd

import (
	"fmt"
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

		err := db.AddTask(todoTask)

		if err != nil {
			panic(err)
		}

		fmt.Println("Task was added succesfully.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
