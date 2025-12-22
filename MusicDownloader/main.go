package main

import (
	"os"
	"path"

	logic "MusicDownloaderCLI/pkg/YTAPI"
)

func main() {
	// cmd.Execute()

	songs := []string{
		"Shape of You Ed Sheeran",
		"Blinding Lights The Weeknd",
		"Uptown Funk Mark Ronson",
		"Rolling in the Deep Adele",
		"Bad Guy Billie Eilish",
		"Someone Like You Adele",
		"Thinking Out Loud Ed Sheeran",
		"Dance Monkey Tones and I",
		"Sunflower Post Malone",
		"Old Town Road Lil Nas X",
		"Perfect Ed Sheeran",
		"Stay Justin Bieber",
		"Levitating Dua Lipa",
		"Believer Imagine Dragons",
		"Senorita Shawn Mendes",
		"Shallow Lady Gaga",
		"Despacito Luis Fonsi",
		"Rockstar Post Malone",
		"Havana Camila Cabello",
		"Closer The Chainsmokers",
		"God's Plan Drake",
		"Lucid Dreams Juice WRLD",
		"Girls Like You Maroon 5",
		"Without Me Halsey",
		"Thank U Next Ariana Grande",
		"Savage Love Jason Derulo",
		"Memories Maroon 5",
		"Don't Start Now Dua Lipa",
		"Peaches Justin Bieber",
		"Watermelon Sugar Harry Styles",
	}
	homepath, _ := os.UserHomeDir()
	downloadlocation := path.Join(homepath, "Downloads")

	logic.Main(downloadlocation, songs)
}
