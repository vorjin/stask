package main

import (
	"fmt"
	"os"
	"path/filepath"

	"stask/cmd"
	"stask/db"
)

func main() {
	path, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting user home directory. Err: %v", err)
		os.Exit(1)
	}
	dbPath := filepath.Join(path, "task-manager.db")

	err = db.BoltDBInit(dbPath)
	if err != nil {
		fmt.Printf("Error initialising Bolt database. Err: %v", err)
		os.Exit(1)
	}
	cmd.Execute()

	err = db.CloseBoltDB()
	if err != nil {
		fmt.Printf("Error closing Bolt database. Err: %v", err)
		os.Exit(1)
	}
}
