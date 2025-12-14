/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cli-task-manager/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
