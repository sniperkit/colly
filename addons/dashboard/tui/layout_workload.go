package tui

import (
	"runtime"
	"strconv"
	"sync"
	"time"

	ui "github.com/sasile/termui"
	cfg "github.com/sniperkit/colly/pkg/config"
)

type TermUI struct {
	cfg                      *cfg.Config
	widget_title             ui.GridBufferer
	widget_sys_info          ui.GridBufferer
	widget_server_info       ui.GridBufferer
	widget_post_iops_chart   *ui.LineChart
	widget_put_iops_chart    *ui.LineChart
	widget_get_iops_chart    *ui.LineChart
	widget_logs              *ui.List
	widget_progress          ui.GridBufferer
	widget_request_bar_chart *ui.BarChart
	widget_post_latency      *ui.BarChart
	widget_put_latency       *ui.BarChart
	widget_get_latency       *ui.BarChart
	iops_get_fifo            *Float64Fifo
	iops_post_fifo           *Float64Fifo
	iops_put_fifo            *Float64Fifo
	logs_fifo                *StringsFifo
	statuses                 map[int]uint64
	ch_done                  chan struct{}
	M                        sync.RWMutex
}

type StringsFifo struct {
	string
	Length      int
	Items       []string
	ch_messages chan string
	lock        sync.Mutex
}

func (self *StringsFifo) Init(length int) {
	self.Length = length
	self.Items = make([]string, length)
	self.ch_messages = make(chan string, 100)
	go func() {
		for msg := range self.ch_messages {
			func() {
				self.lock.Lock()
				defer self.lock.Unlock()
				if len(self.Items) < self.Length {
					self.Items = append(self.Items, msg)
				} else {
					self.Items = self.Items[1:]
					self.Items = append(self.Items, msg)
				}
			}()
		}
	}()
}

func (self *StringsFifo) Insert(msg string) {
	self.ch_messages <- msg
}

func (self *StringsFifo) Get() []string {
	self.lock.Lock()
	defer self.lock.Unlock()
	return self.Items
}

type Float64Fifo struct {
	int
	Length int
	index  int
	Items  []float64
}

func (self *Float64Fifo) Init(length int) {
	self.Length = length
	self.index = 0
	self.Items = make([]float64, length)
}

func (self *Float64Fifo) Insert(i float64) {
	if self.index < self.Length {
		self.Items[self.index] = i
		self.index++
	} else {
		self.Items = self.Items[1:]
		self.Items = append(self.Items, i)
	}
}
func (self *Float64Fifo) Get() []float64 {
	return self.Items
}

func (self *TermUI) ui_set_title(x, y, w, h int) ui.GridBufferer {

	// pp.Println(self.cfg)
	// pp.Println(self.cfg.Title)
	ui_titile_par := ui.NewPar("Running " + self.cfg.Title + " : PRESS q TO QUIT")
	ui_titile_par.Height = h
	ui_titile_par.X = x
	ui_titile_par.Y = y
	ui_titile_par.Width = w
	ui_titile_par.TextFgColor = ui.ColorWhite
	ui_titile_par.BorderLabel = "Title"
	ui_titile_par.BorderFg = ui.ColorCyan

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
		ui.Close()
		close(self.ch_done)
		// ch_done <- true
	})

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})

	/*
		ui.Handle("/timer/1s", func(e ui.Event) {
			draw()
		})

		// handle key q pressing
		ui.Handle("/sys/kbd/q", func(ui.Event) {
			// press q to quit
			ui.StopLoop()
		})

		ui.Handle("/sys/kbd/C-x", func(ui.Event) {
			// handle Ctrl + x combination
		})

		ui.Handle("/sys/kbd", func(ui.Event) {
			// handle all other key pressing
		})

		// handle a 1s timer
		ui.Handle("/timer/1s", func(e ui.Event) {
			t := e.Data.(ui.EvtTimer)
			// t is a EvtTimer
			if t.Count%2 ==0 {
				// do something
			}
		})

		// stop when the crawler is done
		go func() {
			select {
			case <-stopTheUI:
				// wait 10 seconds before closing the ui
				time.Sleep(time.Second * 10)
				termui.StopLoop()
			}
		}()

	*/

	ui.Render(ui_titile_par)
	return ui_titile_par
}

func (self *TermUI) ui_set_servers_info(x, y, w, h int) ui.GridBufferer {
	table1 := ui.NewTable()
	var rows [][]string
	rows = append(rows, []string{"Server", "Port", "TLS Mode"})
	table1.Height += 1
	if len(self.cfg.Global.Servers) == 0 {
		rows = append(rows, []string{self.cfg.Global.Server,
			self.cfg.Global.Port,
			strconv.FormatBool(self.cfg.Global.TLSMode)})
		table1.Height += 2
	} else {
		for _, s := range self.cfg.Global.Servers {
			rows = append(rows, []string{s,
				self.cfg.Global.Port,
				strconv.FormatBool(self.cfg.Global.TLSMode)})
			table1.Height += 2
		}
	}

	table1.Rows = rows
	table1.FgColor = ui.ColorWhite
	table1.BgColor = ui.ColorDefault
	table1.Y = y
	table1.X = x
	table1.Width = w
	table1.BorderLabel = "Servers"
	return table1
}

func (self *TermUI) ui_set_system_info(x, y, w, h int) ui.GridBufferer {
	table1 := ui.NewTable()
	var rows [][]string
	var mem_stat runtime.MemStats
	runtime.ReadMemStats(&mem_stat)
	rows = append(rows, []string{"OS", "CPU's", "Memory"})
	rows = append(rows, []string{runtime.GOOS, strconv.Itoa(runtime.NumCPU()), strconv.FormatInt(int64(mem_stat.Sys), 10)})

	table1.Rows = rows
	table1.FgColor = ui.ColorWhite
	table1.BgColor = ui.ColorDefault
	table1.Y = y
	table1.X = x
	table1.Width = w
	table1.Height = h
	table1.BorderLabel = "System Info"
	return table1
}

func (self *TermUI) ui_set_log_list(x, y, w, h int) *ui.List {
	list := ui.NewList()
	list.ItemFgColor = ui.ColorYellow
	list.BorderLabel = "Log"

	list.Height = h
	list.Width = w
	list.Y = y
	list.X = x
	list.Items = self.logs_fifo.Get()
	return list
}

func (self *TermUI) ui_set_requests_bar_chart(x, y, w, h int) *ui.BarChart {
	bc := ui.NewBarChart()
	bc.BarGap = 3
	bc.BarWidth = 8
	data := []int{}
	bclabels := []string{}
	bc.BorderLabel = "Status codes"

	bc.Data = data
	bc.Width = w
	bc.Height = h
	bc.DataLabels = bclabels
	bc.TextColor = ui.ColorGreen
	bc.BarColor = ui.ColorGreen
	bc.NumColor = ui.ColorYellow
	return bc
}

func (self *TermUI) ui_set_put_latency_bar_chart(x, y, w, h int) *ui.BarChart {
	bc := ui.NewBarChart()
	bc.BarGap = 3
	bc.BarWidth = 8
	data := []int{}
	bclabels := []string{}
	bc.BorderLabel = "[PUT] Latency"

	bc.Data = data
	bc.Width = w
	bc.Height = h
	bc.DataLabels = bclabels
	bc.TextColor = ui.ColorGreen
	bc.BarColor = ui.ColorGreen
	bc.NumColor = ui.ColorYellow
	return bc
}

func (self *TermUI) ui_set_post_latency_bar_chart(x, y, w, h int) *ui.BarChart {
	bc := ui.NewBarChart()
	bc.BarGap = 3
	bc.BarWidth = 8
	data := []int{}
	bclabels := []string{}
	bc.BorderLabel = "[POST] Latency"

	bc.Data = data
	bc.Width = w
	bc.Height = h
	bc.DataLabels = bclabels
	bc.TextColor = ui.ColorGreen
	bc.BarColor = ui.ColorGreen
	bc.NumColor = ui.ColorYellow
	return bc
}

func (self *TermUI) ui_set_get_latency_bar_chart(x, y, w, h int) *ui.BarChart {
	bc := ui.NewBarChart()
	bc.BarGap = 3
	bc.BarWidth = 8
	data := []int{}
	bclabels := []string{}
	bc.BorderLabel = "[GET] Latency"

	bc.Data = data
	bc.Width = w
	bc.Height = h
	bc.DataLabels = bclabels
	bc.TextColor = ui.ColorGreen
	bc.BarColor = ui.ColorGreen
	bc.NumColor = ui.ColorYellow
	return bc
}

func (self *TermUI) ui_get_iops(x, y, w, h int) *ui.LineChart {
	lc2 := ui.NewLineChart()
	lc2.BorderLabel = "[GET] IOPS chart"
	lc2.Mode = "braille"

	lc2.Width = w
	lc2.Height = h
	lc2.X = x
	lc2.Y = y
	lc2.AxesColor = ui.ColorWhite
	lc2.LineColor = ui.ColorCyan | ui.AttrBold
	lc2.Data = self.iops_get_fifo.Get()
	return lc2
}

func (self *TermUI) ui_put_iops(x, y, w, h int) *ui.LineChart {
	lc2 := ui.NewLineChart()
	lc2.BorderLabel = "[PUT] IOPS chart"
	lc2.Mode = "braille"

	lc2.Data = self.iops_put_fifo.Get()
	lc2.Width = w
	lc2.Height = h
	lc2.X = x
	lc2.Y = y
	lc2.AxesColor = ui.ColorWhite
	lc2.LineColor = ui.ColorCyan | ui.AttrBold
	return lc2
}

func (self *TermUI) ui_post_iops(x, y, w, h int) *ui.LineChart {
	lc2 := ui.NewLineChart()
	lc2.BorderLabel = "[POST] IOPS chart"
	lc2.Mode = "braille"

	lc2.Data = self.iops_post_fifo.Get()
	lc2.Width = w
	lc2.Height = h
	lc2.X = x
	lc2.Y = y
	lc2.AxesColor = ui.ColorWhite
	lc2.LineColor = ui.ColorCyan | ui.AttrBold
	return lc2
}

func (self *TermUI) Update_requests(duration time.Duration, post_count, put_count, get_count uint64) {
	seconds := uint64(duration.Seconds())
	if seconds == 0 {
		seconds = 1
	}
	get_iops := get_count / seconds
	post_iops := post_count / seconds
	put_iops := put_count / seconds
	if get_iops > 0 {
		self.iops_get_fifo.Insert(float64(get_iops) / 1000)
	}
	if post_iops > 0 {
		self.iops_post_fifo.Insert(float64(post_iops) / 1000)
	}
	if put_iops > 0 {
		self.iops_put_fifo.Insert(float64(put_iops) / 1000)
	}
	self.widget_get_iops_chart.Data = self.iops_get_fifo.Get()
	self.widget_post_iops_chart.Data = self.iops_post_fifo.Get()
	self.widget_put_iops_chart.Data = self.iops_put_fifo.Get()
}

func (self *TermUI) Refresh_log() {
	self.widget_logs.Items = self.logs_fifo.Get()
	ui.Render(self.widget_logs)
}

func (self *TermUI) Update_status_codes(labels []string, values []int) {
	self.widget_request_bar_chart.Data = values
	self.widget_request_bar_chart.DataLabels = labels
}

func (self *TermUI) Update_put_latency_chart(labels []string, values []int) {
	self.widget_put_latency.Data = values
	self.widget_put_latency.DataLabels = labels
}

func (self *TermUI) Update_post_latency_chart(labels []string, values []int) {
	self.widget_post_latency.Data = values
	self.widget_post_latency.DataLabels = labels
}

func (self *TermUI) Update_get_latency_chart(labels []string, values []int) {
	self.widget_get_latency.Data = values
	self.widget_get_latency.DataLabels = labels
}

func Percentage(value, total int) int {
	return value * total / 100
}

func (self *TermUI) Init_term_ui(cfg *cfg.Config) chan struct{} {
	self.cfg = cfg
	self.ch_done = make(chan struct{})

	self.iops_get_fifo = &Float64Fifo{}
	self.iops_get_fifo.Init(150)

	self.iops_put_fifo = &Float64Fifo{}
	self.iops_put_fifo.Init(150)

	self.iops_post_fifo = &Float64Fifo{}
	self.iops_post_fifo.Init(150)

	self.logs_fifo = &StringsFifo{}
	self.logs_fifo.Init(10)

	self.statuses = make(map[int]uint64)

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	term_hight := ui.TermHeight()

	self.widget_title = self.ui_set_title(0, 0, 50, Percentage(7, term_hight))
	self.widget_server_info = self.ui_set_servers_info(0, 0, 0, 0)
	self.widget_sys_info = self.ui_set_system_info(0, 0, 0, self.widget_server_info.GetHeight())

	self.widget_post_iops_chart = self.ui_post_iops(0, 0, 0, Percentage(30, term_hight))
	self.widget_get_iops_chart = self.ui_get_iops(0, 0, 0, Percentage(30, term_hight))
	self.widget_put_iops_chart = self.ui_put_iops(0, 0, 0, Percentage(30, term_hight))

	self.widget_post_latency = self.ui_set_post_latency_bar_chart(0, 0, 0, Percentage(30, term_hight))
	self.widget_get_latency = self.ui_set_get_latency_bar_chart(0, 0, 0, Percentage(30, term_hight))
	self.widget_put_latency = self.ui_set_put_latency_bar_chart(0, 0, 0, Percentage(30, term_hight))

	self.widget_request_bar_chart = self.ui_set_requests_bar_chart(0, 0, 0, Percentage(20, term_hight))
	self.widget_logs = self.ui_set_log_list(0, 0, 0, Percentage(20, term_hight))

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, self.widget_title),
		),
		ui.NewRow(
			ui.NewCol(6, 0, self.widget_sys_info),
			ui.NewCol(6, 0, self.widget_server_info),
		),
		ui.NewRow(
			ui.NewCol(6, 0, self.widget_put_iops_chart),
			ui.NewCol(6, 0, self.widget_put_latency),
		),
		ui.NewRow(
			ui.NewCol(6, 0, self.widget_post_iops_chart),
			ui.NewCol(6, 0, self.widget_post_latency),
		),
		ui.NewRow(
			ui.NewCol(6, 0, self.widget_get_iops_chart),
			ui.NewCol(6, 0, self.widget_get_latency),
		),
		ui.NewRow(
			ui.NewCol(6, 0, self.widget_logs),
			ui.NewCol(6, 0, self.widget_request_bar_chart),
		),
	)

	ui.Body.Align()
	ui.Render(ui.Body)
	go ui.Loop()
	return self.ch_done
}

func (self *TermUI) Render() {
	ui.Render(ui.Body)
}

func (self *TermUI) Terminate_ui() {
	ui.StopLoop()
	ui.Close()
}

func (self *TermUI) Write(p []byte) (n int, err error) {
	if p == nil {
		return 0, nil
	}
	self.logs_fifo.Insert(string(p))
	if self.widget_logs != nil {
		self.widget_logs.Items = self.logs_fifo.Get()
	}
	ui.Render(self.widget_logs)
	return len(p), nil
}
