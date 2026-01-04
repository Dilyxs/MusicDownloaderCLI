/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// DownloadLocationCmd represents the DownloadLocation command
var DownloadLocationCmd = &cobra.Command{
	Use:   "DownloadLocation",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DownloadLocation called")
		newpath := args[0]
		viper.Set("installation", newpath)
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("could not update location properly!")
		}
	},
}

func init() {
	ConfigCmd.AddCommand(DownloadLocationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// DownloadLocationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// DownloadLocationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
