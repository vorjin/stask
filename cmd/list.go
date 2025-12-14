package cmd

import (
	postgres "cli-task-manager/db"
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `A o quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := postgres.ConnectToDatabase()
		if err != nil {
			panic(err)
		}
		defer db.Close()

		rows, err := db.Query(`SELECT * FROM tasks`)
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			var id int
			var name string
			var done bool
			var createdAt string

			err := rows.Scan(&id, &name, &done, &createdAt)
			if err != nil {
				panic(err)
			}

			fmt.Printf("%d. %s \n", id, name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
