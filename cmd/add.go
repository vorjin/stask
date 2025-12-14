package cmd

import (
	postgres "cli-task-manager/db"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := postgres.ConnectToDatabase()
		if err != nil {
			panic(err)
		}
		defer db.Close()

		sqlQuery := `
		INSERT INTO tasks (title)
		VALUES ($1)
		`
		//interpriting arguments into 1 string
		task := strings.Join(args, " ")

		_, err = db.Exec(sqlQuery, task)
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
