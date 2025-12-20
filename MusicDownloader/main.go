package main

import (
	"fmt"
	"log"
	//"MusicDownloaderCLI/cmd"
	fetcher "MusicDownloaderCLI/pkg/YTAPI"
)

func main() {
	// cmd.Execute()
	data, err := fetcher.FetchYoutubeDetails("A-team Ed Sheeran")
	if err != nil {
		log.Fatalf("ran into err: %v\n", err)
	}
	for _, vid := range data {
		fmt.Println(&vid)
	}
}
