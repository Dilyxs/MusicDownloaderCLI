# MusicDownloaderCLI

A command-line tool written in Go for searching and downloading audio from YouTube.

## Features

- Search YouTube videos by song name/artist
- Interactive selection from top 5 search results
- Auto-select mode for quick downloads
- Audio-only downloads (saves as `.m4a`)
- Configurable download location
- Concurrent architecture with worker pools

## Prerequisites

- Go 1.25.5 or later
- YouTube Data API v3 key ([Get one here](https://console.cloud.google.com/apis/credentials))

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/MusicDownloaderCLI.git
   cd MusicDownloaderCLI
   ```

2. Create a `.env` file with your YouTube API key:
   ```
   YT_API_KEY=your_youtube_api_key_here
   ```

3. Build the application:
   ```bash
   go build -o MusicDownloaderCLI
   ```

## Usage

### Download a Song (Interactive)

Search for a song and choose from the results:

```bash
./MusicDownloaderCLI Get "Bohemian Rhapsody Queen"
```

### Download a Song (Auto-Select)

Automatically download the first search result:

```bash
./MusicDownloaderCLI Get -a "Bohemian Rhapsody Queen"
# or
./MusicDownloaderCLI Get --auto "Bohemian Rhapsody Queen"
```

### Configure Download Location

Set a persistent download directory:

```bash
./MusicDownloaderCLI Config DownloadLocation /path/to/music
```

Override download location for a single command:

```bash
./MusicDownloaderCLI -i /custom/path Get "song name"
```

### Check Current Configuration

Display the current download location:

```bash
./MusicDownloaderCLI
```

## Commands

| Command | Description |
|---------|-------------|
| `Get <song>` | Search and download a song |
| `Get -a <song>` | Auto-download first result |
| `Config DownloadLocation <path>` | Set default download path |

## Flags

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--installation` | `-i` | Override download location | `~/Downloads` |
| `--auto` | `-a` | Auto-select first result | `false` |

## Project Structure

```
MusicDownloader/
├── main.go              # Application entry point
├── cmd/                 # CLI commands
│   ├── root.go          # Root command and global flags
│   ├── Get.go           # Download command
│   ├── Config.go        # Configuration command
│   └── DownloadLocation.go
└── pkg/
    └── YTAPI/           # YouTube API integration
        ├── corelogic.go # Core download logic
        └── fetcher.go   # API fetching and audio download
```

## Configuration Files

| File | Purpose |
|------|---------|
| `.env` | YouTube API key (`YT_API_KEY`) |
| `.musicdownloader.json` | User preferences (download location) |

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Viper](https://github.com/spf13/viper) - Configuration management
- [kkdai/youtube](https://github.com/kkdai/youtube) - YouTube downloading

## License

See [LICENSE](LICENSE) file for details.
