/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/BoanergesJunior/go-project-template/cmd/utils"
	"github.com/BoanergesJunior/go-project-template/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "create-project [project-name]",
	Short: "Create a new Go project with a standard layout",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		buildProject(projectName)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func buildProject(projectName string) {
	fmt.Printf("Building project %s...\n", projectName)

	config.BasePath = projectName

	utils.CreateFolders()
}
