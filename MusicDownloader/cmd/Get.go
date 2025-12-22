/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	logic "MusicDownloaderCLI/pkg/YTAPI"

	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "Get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		song := []string{strings.Join(args, " ")}
		automatic, _ := cmd.Flags().GetBool("auto")
		path, _ := cmd.Root().Flags().GetString("installation-location")
		logic.Main(path, song, automatic)
	},
}

func init() {
	rootCmd.AddCommand(GetCmd)
	GetCmd.Flags().BoolP("auto", "a", false, "automatically select the first matching song!")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
