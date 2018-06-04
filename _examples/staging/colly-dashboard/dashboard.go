package main

import (
	metric "github.com/sniperkit/colly/pkg/metric"
	tui "github.com/sniperkit/colly/plugins/app/dashboard/tui/termui"
	// dash "github.com/sniperkit/colly/plugins/app/dashboard"
	// cui "github.com/sniperkit/colly/plugins/app/dashboard/tui/gocui"
	// tvi "github.com/sniperkit/colly/plugins/app/dashboard/tui/tview"
)

var (
	stopTheUI                    chan bool
	stopTheCrawler               chan bool
	allURLsHaveBeenVisited       chan bool
	allStatisticsHaveBeenUpdated chan bool
	// termUI                       *tui.TermUI
)

func initStatsCollector() {
	collectorStats = metric.NewStatsCollector()
}

func initDashboard() {
	collectorStats = metric.NewStatsCollector()

	stopTheUI = make(chan bool)
	collectorResponseMetrics = make(chan metric.Response)

	go func() {
		tui.Dashboard(collectorStats, stopTheUI, stopTheCrawler)

	}()
}

func updateDashboard() {
	go func() {
		for {
			select {
			case <-stopTheCrawler:
				stopTheUI <- true
				return

			case <-stopTheUI:
				stopTheCrawler <- true

			case snapshot := <-collectorResponseMetrics:
				if collectorStats != nil {
					collectorStats.UpdateStatistics(snapshot)
				}
			}
		}
	}()
}
