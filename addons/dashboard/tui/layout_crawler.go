package tui

import (
	"fmt"
	"math"
	"time"

	"github.com/sasile/termui"
)

/*
	Refs:
	- https://github.com/jessfraz/tdash/blob/master/jenkins.go
	-
*/

func Dashboard(stopTheUI, stopTheCrawler chan bool) {
	if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()

	var snapshots []Snapshot

	// termui.UseTheme("helloworld")

	body := termui.NewGrid()
	body.X = 0
	body.Y = 0
	body.BgColor = termui.ThemeAttr("bg")
	body.Width = termui.TermWidth()

	// gaWidget(body)
	// travisWidget(body)
	// jenkinsWidget(body)

	// Calculate the layout.
	// body.Align()

	mimeTypes := termui.NewPar("")
	mimeTypes.Height = 3
	mimeTypes.TextFgColor = termui.ColorWhite
	mimeTypes.BorderLabel = "Mime Types:"
	mimeTypes.BorderFg = termui.ColorCyan
	mimeTypes.Y = 3

	patternsCount := termui.NewPar("")
	patternsCount.Height = 3
	patternsCount.TextFgColor = termui.ColorWhite
	patternsCount.BorderLabel = "Total Patterns:"
	patternsCount.BorderFg = termui.ColorCyan
	patternsCount.Y = 3

	logWindow := termui.NewList()
	logWindow.ItemFgColor = termui.ColorYellow
	logWindow.BorderLabel = "Log:"
	logWindow.Height = 22
	// logWindow.Y = 3

	logParser := termui.NewList()
	logParser.ItemFgColor = termui.ColorYellow
	logParser.BorderLabel = "Patterns:"
	logParser.Height = 22
	// logParser.Y = 3

	totalBytesDownloaded := termui.NewPar("")
	totalBytesDownloaded.Height = 3
	totalBytesDownloaded.TextFgColor = termui.ColorWhite
	totalBytesDownloaded.BorderLabel = "Data downloaded:"
	totalBytesDownloaded.BorderFg = termui.ColorCyan

	totalNumberOfRequests := termui.NewPar("")
	totalNumberOfRequests.Height = 3
	totalNumberOfRequests.TextFgColor = termui.ColorWhite
	totalNumberOfRequests.BorderLabel = "URLs crawled:"
	totalNumberOfRequests.BorderFg = termui.ColorCyan

	requestsPerSecond := termui.NewPar("")
	requestsPerSecond.Height = 3
	requestsPerSecond.TextFgColor = termui.ColorWhite
	requestsPerSecond.BorderLabel = "URLs/second:"
	requestsPerSecond.BorderFg = termui.ColorCyan

	averageResponseTime := termui.NewPar("")
	averageResponseTime.Height = 3
	averageResponseTime.TextFgColor = termui.ColorWhite
	averageResponseTime.BorderLabel = "Average response time:"
	averageResponseTime.BorderFg = termui.ColorCyan

	numberOfWorkers := termui.NewPar("")
	numberOfWorkers.Height = 3
	numberOfWorkers.TextFgColor = termui.ColorWhite
	numberOfWorkers.BorderLabel = "Number of workers:"
	numberOfWorkers.BorderFg = termui.ColorCyan

	averageSizeInBytes := termui.NewPar("")
	averageSizeInBytes.Height = 3
	averageSizeInBytes.TextFgColor = termui.ColorWhite
	averageSizeInBytes.BorderLabel = "Average response size:"
	averageSizeInBytes.BorderFg = termui.ColorCyan

	numberOfErrors := termui.NewPar("")
	numberOfErrors.Height = 3
	numberOfErrors.TextFgColor = termui.ColorWhite
	numberOfErrors.BorderLabel = "Number of 4xx errors:"
	numberOfErrors.BorderFg = termui.ColorCyan

	elapsedTime := termui.NewPar("")
	elapsedTime.Height = 3
	elapsedTime.TextFgColor = termui.ColorWhite
	elapsedTime.BorderLabel = "Elapsed time:"
	elapsedTime.BorderFg = termui.ColorCyan

	// Tests
	strs := []string{"[0] gizak/termui", "[1] editbox.go", "[2] interrupt.go", "[3] keyboard.go", "[4] output.go", "[5] random_out.go", "[6] dashboard.go", "[7] nsf/termbox-go"}
	list := termui.NewList()
	list.Items = strs
	list.ItemFgColor = termui.ColorYellow
	list.BorderLabel = "List"
	list.Height = 7
	list.Width = 25
	list.Y = 4

	g := termui.NewGauge()
	g.Percent = 50
	g.Width = 50
	g.Height = 3
	g.Y = 11
	g.BorderLabel = "Gauge"
	g.BarColor = termui.ColorRed
	g.BorderFg = termui.ColorWhite
	g.BorderLabelFg = termui.ColorCyan

	spark := termui.Sparkline{}
	spark.Height = 1
	spark.Title = "srv 0:"
	spdata := []int{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6, 4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6, 4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6, 4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}
	spark.Data = spdata
	spark.LineColor = termui.ColorCyan
	spark.TitleColor = termui.ColorWhite

	spark1 := termui.Sparkline{}
	spark1.Height = 1
	spark1.Title = "srv 1:"
	spark1.Data = spdata
	spark1.TitleColor = termui.ColorWhite
	spark1.LineColor = termui.ColorRed

	sp := termui.NewSparklines(spark, spark1)
	sp.Width = 25
	sp.Height = 7
	sp.BorderLabel = "Sparkline"
	sp.Y = 4
	sp.X = 25

	sinps := (func() []float64 {
		n := 220
		ps := make([]float64, n)
		for i := range ps {
			ps[i] = 1 + math.Sin(float64(i)/5)
		}
		return ps
	})()

	lc0 := termui.NewLineChart()
	lc0.BorderLabel = "braille-mode Line Chart"
	lc0.Data = sinps
	lc0.Width = 50
	lc0.Height = 12
	lc0.X = 0
	lc0.Y = 0
	lc0.AxesColor = termui.ColorWhite
	lc0.LineColor = termui.ColorGreen | termui.AttrBold

	lc1 := termui.NewLineChart()
	lc1.BorderLabel = "dot-mode Line Chart"
	lc1.Mode = "dot"
	lc1.Data = sinps
	lc1.Width = 26
	lc1.Height = 12
	lc1.X = 51
	lc1.DotStyle = '+'
	lc1.AxesColor = termui.ColorWhite
	lc1.LineColor = termui.ColorYellow | termui.AttrBold

	lc2 := termui.NewLineChart()
	lc2.BorderLabel = "dot-mode Line Chart"
	lc2.Mode = "dot"
	lc2.Data = sinps[4:]
	lc2.Width = 77
	lc2.Height = 16
	lc2.X = 0
	lc2.Y = 12
	lc2.AxesColor = termui.ColorWhite
	lc2.LineColor = termui.ColorCyan | termui.AttrBold

	// Gauge
	g0 := termui.NewGauge()
	g0.Percent = 40
	g0.Width = 50
	g0.Height = 3
	g0.BorderLabel = "Slim Gauge"
	g0.BarColor = termui.ColorRed
	g0.BorderFg = termui.ColorWhite
	g0.BorderLabelFg = termui.ColorCyan

	gg := termui.NewBlock()
	gg.Width = 50
	gg.Height = 5
	gg.Y = 12
	gg.BorderLabel = "TEST"
	gg.Align()

	g2 := termui.NewGauge()
	g2.Percent = 60
	g2.Width = 50
	g2.Height = 3
	g2.PercentColor = termui.ColorBlue
	g2.Y = 3
	g2.BorderLabel = "Slim Gauge"
	g2.BarColor = termui.ColorYellow
	g2.BorderFg = termui.ColorWhite

	g1 := termui.NewGauge()
	g1.Percent = 30
	g1.Width = 50
	g1.Height = 5
	g1.Y = 6
	g1.BorderLabel = "Big Gauge"
	g1.PercentColor = termui.ColorYellow
	g1.BarColor = termui.ColorGreen
	g1.BorderFg = termui.ColorWhite
	g1.BorderLabelFg = termui.ColorMagenta

	g3 := termui.NewGauge()
	g3.Percent = 50
	g3.Width = 50
	g3.Height = 3
	g3.Y = 11
	g3.BorderLabel = "Gauge with custom label"
	g3.Label = "{{percent}}% (100MBs free)"
	g3.LabelAlign = termui.AlignRight

	g4 := termui.NewGauge()
	g4.Percent = 50
	g4.Width = 50
	g4.Height = 3
	g4.Y = 14
	g4.BorderLabel = "Gauge"
	g4.Label = "Gauge with custom highlighted label"
	g4.PercentColor = termui.ColorYellow
	g4.BarColor = termui.ColorGreen
	g4.PercentColorHighlighted = termui.ColorBlack

	draw := func() {
		snapshot := stats.LastSnapshot()

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

		// time since first snapshot
		timeSinceStart := time.Now().Sub(snapshots[0].Timestamp())
		elapsedTime.Text = fmt.Sprintf("%s", timeSinceStart)

		termui.Render(termui.Body)
	}

	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(6, 0, mimeTypes),
			termui.NewCol(6, 0, patternsCount),
		),
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
		/*
			termui.NewRow(
				termui.NewCol(4, 0, lc0),
				termui.NewCol(4, 0, lc1),
				termui.NewCol(4, 0, lc2),
			),
			termui.NewRow(
				termui.NewCol(4, 0, g0),
				termui.NewCol(4, 0, g1),
				termui.NewCol(4, 0, g2),
			),
			termui.NewRow(
				termui.NewCol(6, 0, g3),
				termui.NewCol(6, 0, g4),
			),
		*/
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
		stopTheCrawler <- true
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/C-c", func(termui.Event) {
		stopTheCrawler <- true
		termui.StopLoop()
	})

	termui.Handle("/timer/1s", func(e termui.Event) {
		draw()
	})

	// stop when the crawler is done
	go func() {
		select {
		// case <-pauseTheUI:
		case <-stopTheUI:
			// wait 10 seconds before closing the ui
			time.Sleep(time.Second * 10)
			termui.StopLoop()
		}
	}()

	termui.Loop()
}
