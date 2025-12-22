package main

import (
	"os"
	"path"

	logic "MusicDownloaderCLI/pkg/YTAPI"
)

func main() {
	// cmd.Execute()
	songs := []string{"A team Ed Sheeran"}
	homepath, _ := os.UserHomeDir()
	downloadlocation := path.Join(homepath, "Downloads")

	logic.Main(downloadlocation, songs)
}
