package cmd

import (
	"encoding/binary"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Long:  `List all of your incomplete tasks currently stored in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("This is your tasks:\n")

		db, err := bolt.Open("task-manager.db", 0644, nil)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = db.View(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte("tasks"))
			if err != nil {
				panic(err)
			}

			cursor := bucket.Cursor()

			for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
				id := binary.BigEndian.Uint64(key)
				fmt.Printf("%d. %s\n", id, value)
			}

			return nil
		})

		if err != nil {
			panic(err)
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
