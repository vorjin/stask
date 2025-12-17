/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"stask/cmd"
	"stask/db"
)

func main() {
	var path = "task-manager.db"
	err := db.BoltDBInit(path)
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
