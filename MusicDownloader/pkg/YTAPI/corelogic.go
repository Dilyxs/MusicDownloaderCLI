package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	DownloaderWorkerCount   int = 20
	ErrorChannelWorkerCount int = 3
	JobFetcherWorkerCount       = 20
)

type Job struct {
	UserPrompt      string
	ID              int
	VideoInfo       *[]RelevantVideoData
	Send            chan int
	UserChosenVideo *RelevantVideoData
}

func (j *Job) String() string {
	return fmt.Sprintf("Job{UserPrompt: %q, ID: %d, VideoInfo: %v, UserChosenVideo: %v}", j.UserPrompt, j.ID, j.VideoInfo, j.UserChosenVideo)
}

type ErrorJob struct {
	Job   *Job
	Error error
}
type Hub struct {
	CounterID            int
	Jobs                 map[*Job]bool
	SongGeneratorChannel chan string
	JobFetchedChannel    chan *Job
	DownloadsChannel     chan *Job
	ErrorChannel         chan *ErrorJob
	DownloadLocation     string
	mu                   sync.Mutex
}

func (hub *Hub) GetCounterValue() int {
	hub.mu.Lock()
	id := hub.CounterID
	hub.CounterID++
	hub.mu.Unlock()
	return id
}

func (hub *Hub) AddJob(job *Job) {
	hub.mu.Lock()
	hub.Jobs[job] = true
	defer hub.mu.Unlock()
}

func (hub *Hub) RemoveJob(job *Job) {
	hub.mu.Lock()
	if _, ok := hub.Jobs[job]; ok {
		close(job.Send)
		delete(hub.Jobs, job)
	}
	hub.mu.Unlock()
}

func FetcherWorker(wg *sync.WaitGroup, hub *Hub) {
	defer wg.Done()
	for userintput := range hub.SongGeneratorChannel {
		PossibleSongs, err := FetchYoutubeDetails(userintput)
		if err != nil {
			hub.ErrorChannel <- &ErrorJob{&Job{UserPrompt: userintput}, err}
			return
		}
		newjob := &Job{UserPrompt: userintput, ID: hub.GetCounterValue(), VideoInfo: &PossibleSongs}
		hub.JobFetchedChannel <- newjob
		fmt.Println("JOB HAS BEEN SENT TO MAIN FOR CHOICE!")
	}
}

func GetUserInput(job *Job) {
	valueset := make(map[int]*RelevantVideoData)
	fmt.Println("Pick the song which you want!")

	for index, video := range *job.VideoInfo {
		fmt.Printf("%d : %v\n", index, video.Title)
		valueset[index] = &video
	}

	r := bufio.NewReader(os.Stdin)
	var IsOkChosenVideo bool

	for !IsOkChosenVideo {
		fmt.Println("\npls pick the Video Title you find appropriate! ")

		input, _ := r.ReadString('\n')
		userinput := strings.TrimSpace(input)
		index, err := strconv.Atoi(userinput)

		if err != nil {
			fmt.Println("pls make sure to pick an integer!")
		} else {
			if video, ok := valueset[index]; ok {
				job.UserChosenVideo = video
				IsOkChosenVideo = true
			} else {
				fmt.Println("Make sure that chosen integer is a possible choice!")
			}
		}
	}
}

func DownloaderWorker(hub *Hub, wg *sync.WaitGroup, downloadDirectoryPath string) {
	defer wg.Done()
	for download := range hub.DownloadsChannel {
		err := DownloadAudio(download.UserChosenVideo.VideoID, download.UserChosenVideo.Title)
		if err != nil {
			hub.ErrorChannel <- &ErrorJob{download, err}
			return
		}
		hub.RemoveJob(download)
		fmt.Printf("SONG DOWNLOADED! at %v", downloadDirectoryPath)
	}
}

func ErrorWorker(hub *Hub, wg *sync.WaitGroup) {
	defer wg.Done()
	for error := range hub.ErrorChannel {
		fmt.Printf("Unfortunately, could not download song for : %v as this happened: %v\n", error.Job.UserChosenVideo.Title, error.Error)
		hub.RemoveJob(error.Job)
	}
}

func Main(DownloadLocation string, songs []string) {
	hub := Hub{
		CounterID:            0,
		Jobs:                 make(map[*Job]bool),
		SongGeneratorChannel: make(chan string, JobFetcherWorkerCount),
		JobFetchedChannel:    make(chan *Job, 1),
		DownloadsChannel:     make(chan *Job, DownloaderWorkerCount),
		ErrorChannel:         make(chan *ErrorJob, ErrorChannelWorkerCount),
		DownloadLocation:     DownloadLocation,

		mu: sync.Mutex{},
	}
	var SongFetcherWaitgroup sync.WaitGroup
	var DownloadWorkerWaitgroup sync.WaitGroup
	var ErrorChannelWaitgroup sync.WaitGroup

	// lauch ErrorWorkers
	for i := range ErrorChannelWorkerCount {
		_ = i
		ErrorChannelWaitgroup.Add(1)
		go ErrorWorker(&hub, &ErrorChannelWaitgroup)
	}

	// lauch DownloadWorkers to get ready for the download on the spot
	for i := range DownloaderWorkerCount {
		_ = i
		DownloadWorkerWaitgroup.Add(1)
		go DownloaderWorker(&hub, &DownloadWorkerWaitgroup, DownloadLocation)
	}
	// launch FetcherWorker to get ready to fetch data which will then be processed by main to pick a song!
	for i := range JobFetcherWorkerCount {
		_ = i
		SongFetcherWaitgroup.Add(1)
		go FetcherWorker(&SongFetcherWaitgroup, &hub)
	}
	// cleanup function
	go func() {
		SongFetcherWaitgroup.Wait()
		close(hub.JobFetchedChannel)
		fmt.Println("Fetcher boys are done!")
	}()

	// Start the Engine by pushing each song into the SongFetcherWorker's channel
	for _, song := range songs {
		hub.SongGeneratorChannel <- song
	}
	// struggling to send songs
	close(hub.SongGeneratorChannel)

	for job := range hub.JobFetchedChannel {
		// Ask user for preferred video
		GetUserInput(job)
		hub.DownloadsChannel <- job
	}
	close(hub.DownloadsChannel)

	DownloadWorkerWaitgroup.Wait()
	fmt.Println("done waiting for the downloads bros!")
	close(hub.ErrorChannel)
	fmt.Println("Error Channel closed!")
	ErrorChannelWaitgroup.Wait()
}
