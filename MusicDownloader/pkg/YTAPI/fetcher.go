package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/subosito/gotenv"
)

type YoutubeID struct {
	Kind    string `json:"kind"`
	VideoID string `json:"videoId"`
}
type YoutubeThubnmailFormat struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type YoutubeThumbnails struct {
	Default YoutubeThubnmailFormat `json:"default"`
	Medium  YoutubeThubnmailFormat `json:"medium"`
	High    YoutubeThubnmailFormat `json:"high"`
}
type YoutubeVideoDetails struct {
	PublishedData        time.Time         `json:"publishedAt"`
	ChannelID            string            `json:"channelId"`
	Title                string            `json:"title"`
	Description          string            `json:"description"`
	Thumbnails           YoutubeThumbnails `json:"thumbnails"`
	ChannelTitle         string            `json:"channelTitle"`
	LiveBroadcastContent string            `json:"liveBroadcastContent"`
	PublishTime          time.Time         `json:"publishTime"` // not sure why there are 2 time, some goofy shit with yt api
}
type YoutubeItem struct {
	Kind         string              `json:"kind"`
	Etag         string              `json:"etag"`
	ID           YoutubeID           `json:"id"`
	VideoDetails YoutubeVideoDetails `json:"snippet"`
}

type PageDetails struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}
type YoutubeAPIResponse struct {
	TypeOfAPIResponse string        `json:"kind"`
	Etag              string        `json:"etag"`
	NextPageToken     string        `json:"nextPageToken"`
	RegionCode        string        `json:"regionCode"`
	PageInfo          PageDetails   `json:"pageInfo"`
	YoutubeVids       []YoutubeItem `json:"items"`
}
type RelevantVideoData struct {
	Title        string
	Description  string
	ChannelTitle string
	VideoID      string
}

func (vid *RelevantVideoData) String() string {
	return fmt.Sprintf("Video Title: %v, Description: %v, ChannelTitle: %v\n", vid.Title, vid.Description, vid.ChannelTitle)
}

const (
	YoutubePart string = "snippet"
	YoutubeLink string = "https://www.googleapis.com/youtube/v3/search"
)

func LoadVars(path string) error {
	if path == "" {
		return gotenv.Load()
	}
	return gotenv.Load(path)
}

func FetchYoutubeDetails(userscontent string) ([]RelevantVideoData, error) {
	if err := LoadVars(""); err != nil {
		return []RelevantVideoData{}, err
	}
	apikey := os.Getenv("YT_API_KEY")
	if len(apikey) == 0 {
		return []RelevantVideoData{}, fmt.Errorf("cannot find api key(.env) file: %v", apikey)
	}
	u, _ := url.Parse(YoutubeLink)
	params := url.Values{}
	params.Add("part", YoutubePart)
	params.Add("q", userscontent)
	params.Add("key", apikey)
	params.Add("maxResults", "25")
	u.RawQuery = params.Encode()
	resp, err := http.Get(u.String())
	if resp.StatusCode != 200 || err != nil {
		return []RelevantVideoData{}, fmt.Errorf("request for api did not work: %v", resp.StatusCode)
	}
	var Response YoutubeAPIResponse
	newerr := json.NewDecoder(resp.Body).Decode(&Response)
	if newerr != nil {
		return []RelevantVideoData{}, fmt.Errorf("could not decode YoutubeResponse Response properly: %v", newerr)
	}

	// now deal with cleaning up the struct to get data users cares about!
	RelevantVideos := make([]RelevantVideoData, len(Response.YoutubeVids))
	for i, vid := range Response.YoutubeVids {
		RelevantVideos[i] = RelevantVideoData{Title: vid.VideoDetails.Title, Description: vid.VideoDetails.Description, ChannelTitle: vid.VideoDetails.ChannelTitle, VideoID: vid.ID.VideoID}
	}
	return RelevantVideos, nil
}
