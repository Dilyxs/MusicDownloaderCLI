/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "MusicDownloaderCLI",
	Short: "A ClI tool to download all your favorites musics!",
	Long: `Basically give the song name and author name, and it will return possible options from which you
	can choose to download from. Concurrency coming soon!`,
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("installation")
		fmt.Printf("current installation-location is : %s", path)
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, "Downloads")
	rootCmd.PersistentFlags().StringP("installation", "i", path, "Define the default download location of your music videos")
	viper.BindPFlag("installation", rootCmd.PersistentFlags().Lookup("installation"))
}
