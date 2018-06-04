package tui

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/sasile/termui"

	"github.com/sniperkit/colly/pkg/metric"
)

var (
	errorMessages []string
	stats         *metric.Statistics
)

type TermUI struct {
	isReady                                   bool
	widgetTitle                               termui.GridBufferer
	widgetSysInfo                             termui.GridBufferer
	widgetProgress                            termui.GridBufferer
	widgetLogs                                *termui.List
	widgetWatcher                             *termui.List
	widgetTopContentTypes                     *termui.List
	widgetTopStatusCodes                      *termui.List
	widgetTopWhitelistFilters_URL             *termui.List
	widgetTopBlacklistFilters_URL             *termui.List
	widgetTopWhitelistFilters_Domain          *termui.List
	widgetTopBlacklistFilters_Domain          *termui.List
	widgetTopWhitelistFilters_ResponseHeaders *termui.List
	widgetTopBlacklistFilters_ResponseHeaders *termui.List
	widgetTopWhitelistFilters_ResponseBody    *termui.List
	widgetTopBlacklistFilters_ResponseBody    *termui.List
	widgetElapsedTime                         *termui.Par
	widgetTotalBytesDownloaded                *termui.Par
	widgetTotalNumberOfRequests               *termui.Par
	widgetRequestsPerSecond                   *termui.Par
	widgetAverageResponseTime                 *termui.Par
	widgetNumberOfWorkers                     *termui.Par
	widgetAverageSizeInBytes                  *termui.Par
	widgetPostRequestLatency                  *termui.BarChart
	widgetPutRequestLatency                   *termui.BarChart
	widgetGetRequestLatency                   *termui.BarChart
	done                                      chan struct{}
	lock                                      sync.RWMutex
}

func Dashboard(stats *metric.Statistics, stopTheUI, stopTheCrawler chan bool) {
	if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()

	var snapshots []metric.Snapshot
	// stats = collectorMetrics.NewStatsCollector()

	body := termui.NewGrid()
	body.X = 0
	body.Y = 0
	body.BgColor = termui.ThemeAttr("bg")
	body.Width = termui.TermWidth()

	logWindow := termui.NewList()
	logWindow.ItemFgColor = termui.ColorYellow
	logWindow.BorderLabel = "Logs"
	logWindow.Height = 22

	logParser := termui.NewList()
	logParser.ItemFgColor = termui.ColorYellow
	logParser.BorderLabel = "Parser"
	logParser.Height = 22

	totalBytesDownloaded := termui.NewPar("")
	totalBytesDownloaded.Height = 3
	totalBytesDownloaded.TextFgColor = termui.ColorWhite
	totalBytesDownloaded.BorderLabel = "Data downloaded"
	totalBytesDownloaded.BorderFg = termui.ColorCyan

	totalNumberOfRequests := termui.NewPar("")
	totalNumberOfRequests.Height = 3
	totalNumberOfRequests.TextFgColor = termui.ColorWhite
	totalNumberOfRequests.BorderLabel = "URLs crawled"
	totalNumberOfRequests.BorderFg = termui.ColorCyan

	requestsPerSecond := termui.NewPar("")
	requestsPerSecond.Height = 3
	requestsPerSecond.TextFgColor = termui.ColorWhite
	requestsPerSecond.BorderLabel = "URLs/second"
	requestsPerSecond.BorderFg = termui.ColorCyan

	averageResponseTime := termui.NewPar("")
	averageResponseTime.Height = 3
	averageResponseTime.TextFgColor = termui.ColorWhite
	averageResponseTime.BorderLabel = "Average response time"
	averageResponseTime.BorderFg = termui.ColorCyan

	numberOfWorkers := termui.NewPar("")
	numberOfWorkers.Height = 3
	numberOfWorkers.TextFgColor = termui.ColorWhite
	numberOfWorkers.BorderLabel = "Number of workers"
	numberOfWorkers.BorderFg = termui.ColorCyan

	averageSizeInBytes := termui.NewPar("")
	averageSizeInBytes.Height = 3
	averageSizeInBytes.TextFgColor = termui.ColorWhite
	averageSizeInBytes.BorderLabel = "Average response size"
	averageSizeInBytes.BorderFg = termui.ColorCyan

	numberOfErrors := termui.NewPar("")
	numberOfErrors.Height = 3
	numberOfErrors.TextFgColor = termui.ColorWhite
	numberOfErrors.BorderLabel = "Number of 4xx errors"
	numberOfErrors.BorderFg = termui.ColorCyan

	topRequestsContentTypes := termui.NewList()
	topRequestsContentTypes.ItemFgColor = termui.ColorYellow
	topRequestsContentTypes.BorderLabel = "Top10 Resp. Content-Types"
	topRequestsContentTypes.Height = 12

	topRequestsStatusCodes := termui.NewList()
	topRequestsStatusCodes.ItemFgColor = termui.ColorYellow
	topRequestsStatusCodes.BorderLabel = "Top10 Resp. Status Codes"
	topRequestsStatusCodes.Height = 12

	topRequestsWhitelistMatches := termui.NewList()
	topRequestsWhitelistMatches.ItemFgColor = termui.ColorYellow
	topRequestsWhitelistMatches.BorderLabel = "Top10 Whitelist match(es)"
	topRequestsWhitelistMatches.Height = 12

	topRequestsBlacklistMatches := termui.NewList()
	topRequestsBlacklistMatches.ItemFgColor = termui.ColorYellow
	topRequestsBlacklistMatches.BorderLabel = "Top10 Blacklist match(es)"
	topRequestsBlacklistMatches.Height = 12

	elapsedTime := termui.NewPar("")
	elapsedTime.Height = 3
	elapsedTime.TextFgColor = termui.ColorWhite
	elapsedTime.BorderLabel = "Elapsed time:"
	elapsedTime.BorderFg = termui.ColorCyan

	draw := func() {
		snapshot := stats.LastSnapshot()
		// log.Println("snapshot.Timestamp()=", snapshot.Timestamp().String(), "len(snapshots)=", len(snapshots))

		// ignore empty updates
		if snapshot.Timestamp().IsZero() {
			return
		}

		// don't update if there is no new snapshot available
		if len(snapshots) > 0 {
			previousSnapShot := snapshots[len(snapshots)-1]
			if snapshot.Timestamp() == previousSnapShot.Timestamp() {
				return
			}
		}

		// capture the latest snapshot
		snapshots = append(snapshots, snapshot)

		// log messages
		logWindow.Items = stats.LastLogMessages(20)

		// total number of requests
		totalNumberOfRequests.Text = fmt.Sprintf("%d", snapshot.TotalNumberOfRequests())

		// total number of bytes downloaded
		totalBytesDownloaded.Text = formatBytes(snapshot.TotalSizeInBytes())

		// requests per second
		requestsPerSecond.Text = fmt.Sprintf("%.1f", snapshot.RequestsPerSecond())

		// average response time
		averageResponseTime.Text = fmt.Sprintf("%s", snapshot.AverageResponseTime())

		// number of workers
		numberOfWorkers.Text = fmt.Sprintf("%d", snapshot.NumberOfWorkers())

		// average request size
		averageSizeInBytes.Text = formatBytes(snapshot.AverageSizeInBytes())

		// number of errors
		numberOfErrors.Text = fmt.Sprintf("%d", snapshot.NumberOfErrors())

		// list of responses' content types
		topRequestsContentTypes.Items = stats.TopContentTypes(10)

		// list of responses' status code
		topRequestsStatusCodes.Items = stats.TopStatusCodes(10)

		// list of top black listed pattern matches
		topRequestsBlacklistMatches.Items = stats.TopFiltersMatches(10)

		// list of top black listed pattern matches
		topRequestsWhitelistMatches.Items = stats.TopFiltersMatches(10)

		// time since first snapshot
		timeSinceStart := time.Now().Sub(snapshots[0].Timestamp())
		elapsedTime.Text = fmt.Sprintf("%s", timeSinceStart)

		termui.Render(termui.Body)
	}

	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(7, 0, logWindow),
			termui.NewCol(5, 0, logParser),
		),
		termui.NewRow(
			termui.NewCol(3, 0, totalBytesDownloaded),
			termui.NewCol(3, 0, totalNumberOfRequests),
			termui.NewCol(3, 0, requestsPerSecond),
			termui.NewCol(3, 0, averageResponseTime),
		),
		termui.NewRow(
			termui.NewCol(3, 0, numberOfWorkers),
			termui.NewCol(3, 0, numberOfErrors),
			termui.NewCol(3, 0, averageSizeInBytes),
			termui.NewCol(3, 0, elapsedTime),
		),
		termui.NewRow(
			termui.NewCol(3, 0, topRequestsContentTypes),
			termui.NewCol(3, 0, topRequestsStatusCodes),
			termui.NewCol(3, 0, topRequestsWhitelistMatches),
			termui.NewCol(3, 0, topRequestsBlacklistMatches),
		),
	)

	termui.Body.Align()
	termui.Render(termui.Body)

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		termui.Body.Width = termui.TermWidth()
		termui.Body.Align()
		termui.Clear()
		termui.Render(termui.Body)
	})

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		// log.Println("touch 'q' is pressed")
		// stopTheUI <- true
		// stopTheCrawler <- true
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/C-c", func(termui.Event) {
		// log.Println("touches 'ctrl+c' are pressed")
		// stopTheUI <- true
		// stopTheCrawler <- true
		termui.StopLoop()
	})

	termui.Handle("/timer/1s", func(e termui.Event) {
		draw()
	})

	// stop when the crawler is done
	go func() {
		select {
		// case <-pauseTheUI:
		// case <-stopTheCrawler:

		case <-stopTheUI:
			// log.Println("stopTheUI event")
			// wait 10 seconds before closing the ui
			time.Sleep(time.Second * 2)
			termui.StopLoop()
			os.Exit(1)
		}
	}()

	termui.Loop()
}
