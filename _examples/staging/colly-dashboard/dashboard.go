package main

import (
	// triiger "github.com/sadlil/go-trigger"

	metric "github.com/sniperkit/colly/pkg/metric"
	tui "github.com/sniperkit/colly/plugins/app/dashboard/tui/termui"
	// dash "github.com/sniperkit/colly/plugins/app/dashboard"
	// cui "github.com/sniperkit/colly/plugins/app/dashboard/tui/gocui"
	// tvi "github.com/sniperkit/colly/plugins/app/dashboard/tui/tview"
)

var (
	enableUI                     bool = false
	stopTheUI                    chan bool
	stopTheCrawler               chan bool
	allURLsHaveBeenVisited       chan bool
	allStatisticsHaveBeenUpdated chan bool
	// termUI                       *tui.TermUI
)

func initStatsCollector() {
	collectorStats = metric.NewStatsCollector()
	// collectorStats = metric.NewStatsCollector(stopTheUI, stopTheCrawler)
}

func initDashboard() {
	stopTheUI = make(chan bool)
	collectorResponseMetrics = make(chan metric.Response)
	go func() {
		tui.Dashboard(stopTheUI, stopTheCrawler)
	}()
}

func updateDashboard() {
	go func() {
		for {
			select {
			//case <-allURLsHaveBeenVisited:
			//	allStatisticsHaveBeenUpdated <- true
			//	return

			case <-stopTheCrawler:
				stopTheUI <- true

			case snapshot := <-collectorResponseMetrics:
				if collectorStats != nil {
					collectorStats.UpdateStatistics(snapshot)
				}
			}
		}
	}()
}
