package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"stask/db"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clears TODO tasks list",
	Run: func(cmd *cobra.Command, args []string) {
		err := db.DeleteTasksBucket()
		if err != nil {
			panic(err)
		}

		color.Green("TODO list was succesfully cleared!")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
