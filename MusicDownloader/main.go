package main

import (
	"fmt"

	"MusicDownloaderCLI/cmd"

	"github.com/spf13/viper"
)

func init() {
	// look for config file where the user is startin it!
	viper.AddConfigPath("./")
	viper.SetConfigName(".musicdownloader")
	viper.SetConfigType("json")
	err := viper.ReadInConfig() // if any error, we just ignore it
	if err != nil {
		viper.SafeWriteConfig()
		fmt.Println("made a new config file")
	}
}

func main() {
	cmd.Execute()
}
