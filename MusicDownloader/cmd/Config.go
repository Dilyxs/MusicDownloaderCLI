/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ConfigCmd represents the Config command
var ConfigCmd = &cobra.Command{
	Use:   "Config",
	Short: "Configure the settings for your CLI tool",
	Long: `This will create a .json file with your preffered settings, 
	for now this includes only the downloadLocation`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Config called")
	},
}

func init() {
	rootCmd.AddCommand(ConfigCmd)
}
