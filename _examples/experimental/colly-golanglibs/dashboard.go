package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/asdine/storm"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	// core
	colly "github.com/sniperkit/colly/pkg"

	// experimental addons
	tui "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/termui"
	tui_hist "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/termui/histogram"
	// dash "github.com/sniperkit/colly/plugins/cmd/dashboard"
	// cui "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/gocui"
	// tvi "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/tview"
)

/*
	Refs:
	- https://github.com/c1982/mcap/blob/master/main.go
*/

var (
	wl_id                int32 = -1
	enable_ui            bool  = true
	dataBfr              []byte
	ex_group             sync.WaitGroup
	ch_resp_latency      chan time.Duration
	ch_req_latency       chan time.Duration
	ch_post_latency      chan time.Duration
	ch_put_latency       chan time.Duration
	ch_get_latency       chan time.Duration
	ch_statuses          chan int
	term_ui              *tui.TermUI
	LatencyCollectorGet  tui_hist.LatencyHistogram // tui.LatencyCollector
	LatencyCollectorPut  tui_hist.LatencyHistogram //tui.LatencyCollector
	LatencyCollectorPost tui_hist.LatencyHistogram //tui.LatencyCollector
	StatusesCollector    tui.CollectorStatus
	uiWaitGroup          = &sync.WaitGroup{}
	stopTheUI            chan bool
)

func dashboardMcap() {
	var err error
	conn, err = storm.Open(databaseName)
	if err != nil {
		panic(err)
	}
	conn.Init(data{})
	var exchanges = []string{
		"koineks",
		"btcturk",
		"ovis",
		"paribu",
		"vebitcoin",
		"koinim",
		"bithesap",
		"sistemkoin",
	}

	app := tview.NewApplication()
	cols := tview.NewTable().SetSeparator(tview.GraphicsVertBar)
	cols.SetBorder(true).SetTitle("Turkish Cryptocurrency Market Capitalizations").SetTitleColor(tcell.ColorWhite)
	loadData(cols, exchanges)

	err = app.SetRoot(cols, true).Run()
	if err != nil {
		panic(err)
	}

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func get_workload_id() int {
	return int(atomic.AddInt32(&wl_id, 1))
}

func generate_executors(term_ui *tui.TermUI) {
	ch_post_latency = LatencyCollectorPost.New()
	ch_put_latency = LatencyCollectorPut.New()
	ch_get_latency = LatencyCollectorGet.New()
	ch_statuses = StatusesCollector.New(160, 1)

	/*
		for Name, workload := range collyConfig.Workloads {
			log.Println("Adding executor for ", Name)
			workload.Id = get_workload_id()
			e := &httpblaster.Executor{
				Globals:         collyConfig.Global,
				Workload:        workload,
				Host:            collyConfig.Global.Server,
				Hosts:           collyConfig.Global.Servers,
				TLS_mode:        collyConfig.Global.TLSMode,
				Data_bfr:        dataBfr,
				TermUi:          term_ui,
				Ch_get_latency:  ch_get_latency,
				Ch_put_latency:  ch_put_latency,
				Ch_post_latency: ch_post_latency,
				Ch_statuses:     ch_statuses,
				DumpFailures:    dump_failures,
				DumpLocation:    dump_location}
			executors = append(executors, e)
		}
	*/

}

func dump_latencies_histograms() {
	latency_get := make(map[int64]int)
	latency_post := make(map[int64]int)
	latency_put := make(map[int64]int)
	total_get := 0
	total_put := 0
	total_post := 0

	/*
		for _, e := range executors {
			hist := e.LatencyHist()
			if e.GetType() == "GET" {
				for k, v := range hist {
					latency_get[k] += v
					total_get += v
				}
			} else if e.GetType() == "POST" {
				for k, v := range hist {
					latency_post[k] += v
					total_post += v
				}
			} else {
				for k, v := range hist {
					latency_put[k] += v
					total_put += v
				}
			}
		}
	*/

	dump_latency_histogram(latency_get, total_get, "GET")
	dump_latency_histogram(latency_put, total_put, "PUT")
	dump_latency_histogram(latency_post, total_post, "POST")
}

func remap_latency_histogram(hist map[int64]int) map[int64]int {
	res := make(map[int64]int)
	for k, v := range hist {
		if k > 10000 { //1 sec
			res[10000] += v
		} else if k > 5000 { //500 mili
			res[5000] += v
		} else if k > 1000 { // 100mili
			res[1000] += v
		} else if k > 100 { //10 mili
			res[100] += v
		} else if k > 50 { //5 mili
			res[50] += v
		} else if k > 20 { //2 mili
			res[20] += v
		} else if k > 10 { //1 mili
			res[10] += v
		} else { //below 1 mili
			res[k] += v
		}
	}
	return res
}

/*
func start_executors() {
	ex_group.Add(len(executors))
	start_time = time.Now()
	for _, e := range executors {
		e.Start(&ex_group)
	}
}
*/

func wait_for_completion() {
	log.Println("Wait for executors to finish")
	ex_group.Wait()
	end_time = time.Now()
	close(ch_get_latency)
	close(ch_put_latency)
	close(ch_post_latency)
}

func wait_for_ui_completion(ch_done chan struct{}) {
	if enable_ui {
		select {
		case <-ch_done:
			break
		case <-time.After(time.Second * 10):
			close(ch_done)
			break
		}
	}
}

func enable_tui() chan struct{} {
	if enable_ui {
		// pp.Println("collyConfig: ", collyConfig)
		// log.Fatal("test collyConfig")
		term_ui = &tui.TermUI{}
		ch_done := term_ui.Init_term_ui(collyConfig)
		go func() {
			defer term_ui.Terminate_ui()
			tick := time.Tick(time.Millisecond * 500)
			for {
				select {
				case <-ch_done:
					return
				case <-tick:
					// term_ui.Update_get_latency_chart(LatencyCollectorGet.Get())
					// term_ui.Update_post_latency_chart(LatencyCollectorPost.Get())
					// term_ui.Update_put_latency_chart(LatencyCollectorPut.Get())
					// term_ui.Update_status_codes(StatusesCollector.Get())
					term_ui.Refresh_log()
					term_ui.Render()
				}
			}
		}()
		return ch_done
	}
	return nil
}

func dump_latency_histogram(histogram map[int64]int, total int, req_type string) ([]string, []float64) {
	var keys []int
	var prefix string
	title := "type \t usec \t\t percentage\n"
	if req_type == "GET" {
		prefix = "GetHist"
	} else {
		prefix = "PutHist"
	}
	strout := fmt.Sprintf("%s Latency Histograms:\n", prefix)
	hist := remap_latency_histogram(histogram)
	for k := range hist {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	log.Debugln("latency hist wait released")
	res_strings := []string{}
	res_values := []float64{}

	for _, k := range keys {
		v := hist[int64(k)]
		res_strings = append(res_strings, fmt.Sprintf("%5d", k*100))
		value := float64(v*100) / float64(total)
		res_values = append(res_values, value)
	}

	if len(res_strings) > 0 {
		strout += title
		for i, v := range res_strings {
			strout += fmt.Sprintf("%s: %s \t\t %3.4f%%\n", prefix, v, res_values[i])
		}
	}
	log.Println(strout)
	return res_strings, res_values
}

func loadData(t *tview.Table, exchanges []string) {

	t.SetCell(0, 0, tview.NewTableCell("exchange").SetTextColor(tcell.ColorYellow))
	t.SetCell(0, 1, tview.NewTableCell("capital").SetTextColor(tcell.ColorYellow))

	list := dataList(exchanges)
	for i := 0; i < len(list); i++ {
		t.SetCell(i+1, 0, tview.NewTableCell(list.ToMarket(i)).SetTextColor(tcell.ColorDarkCyan))
		t.SetCell(i+1, 1, tview.NewTableCell(list.ToCap(i)).SetAlign(tview.AlignRight))
	}

	t.SetCell(len(list)+1, 0, tview.NewTableCell(""))
	t.SetCell(len(list)+1, 1, tview.NewTableCell(""))
	t.SetCell(len(list)+2, 0, tview.NewTableCell("total ").SetAlign(tview.AlignRight).SetTextColor(tcell.ColorDarkRed))
	t.SetCell(len(list)+2, 1, tview.NewTableCell(list.Total()))
}

func dataList(exchanges []string) (list capdata) {
	c := colly.NewCollector()
	c.DisableCookies()
	for i := 0; i < len(exchanges); i++ {
		d := data{}
		name := exchanges[i]
		cap, capstr, err := getCapital(c, name)
		d.Market = name
		d.SizeStr = capstr
		if err != nil {
			d.Size = -1
		} else {
			d.Size = cap
		}
		err = conn.Save(d)
		if err != nil {
			log.Println("save error:", err)
		}
		list = append(list, d)
	}

	sort.Sort(list)
	return
}

func getCapital(c *colly.Collector, exchange string) (cap float64, capStr string, err error) {
	var currencyValue = ""
	c.OnHTML("span[data-currency-value]", func(e *colly.HTMLElement) {
		capStr = "   " + e.Text
		currencyValue = strings.TrimPrefix(e.Text, "$")
		currencyValue = strings.Replace(currencyValue, ",", "", -1)
	})
	err = c.Visit("https://coinmarketcap.com/exchanges/" + exchange + "/")
	if err != nil {
		return cap, capStr, err
	}
	cap, err = strconv.ParseFloat(currencyValue, 32)
	return cap, capStr, err
}
