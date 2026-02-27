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
	dbDir := filepath.Join(path, ".stask")
	err = os.MkdirAll(dbDir, 0o700)
	if err != nil {
		fmt.Printf("Error creating .stask directory. Err: %v", err)
		os.Exit(1)
	}
	dbPath := filepath.Join(dbDir, "task-manager.db")

	taskStore, err := db.NewBoltTaskStore(dbPath, "tasks")
	if err != nil {
		fmt.Printf("Error initialising Bolt database. Err: %v", err)
		os.Exit(1)
	}

	defer func() {
		err := taskStore.Close()
		if err != nil {
			fmt.Printf("Error closing Bolt database. Err: %v", err)
			os.Exit(1)
		}
	}()

	cmd.Execute(taskStore)
}
