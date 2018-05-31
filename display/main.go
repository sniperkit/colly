package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Jeffail/gabs"
	ui "github.com/gizak/termui"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	log "github.com/sirupsen/logrus"
)

func main() {
	u := NewUI()
	go u.Connect()

	if err := ui.Init(); err != nil {
		panic(err)
	}

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/c", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Loop()
	defer ui.Close()
}

type UI struct {
	Table *ui.Table
	Bars  *ui.BarChart
}

func NewUI() *UI {
	table := ui.NewTable()
	table.BorderLabel = "Top Logs"
	table.Height = 100
	table.Width = 50
	table.Y = 0

	bars := ui.NewBarChart()
	bars.BorderLabel = "Level Counts"
	bars.Width = 50
	bars.Height = 10
	bars.BarWidth = 5
	bars.X = 50

	return &UI{
		Table: table,
		Bars:  bars,
	}
}

func (u *UI) Connect() {
	t := time.Now().Add(time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), t)
	defer cancel()
	conn, _, err := ws.Dial(ctx, "ws://localhost:6002/ws", nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		data, _, err := wsutil.ReadServerData(conn)
		if err != nil {
			log.WithField("err", err).Fatal("Could not read server data.")
			return
		}
		container, err := gabs.ParseJSON(data)
		if err != nil {
			log.WithField("err", err).Fatal("Could not unmarshal server JSON.")
			return
		}

		u.update(container)
	}
}

func (u *UI) update(container *gabs.Container) {
	u.updateTable(container)
	u.updateBars(container)
}

func (u *UI) updateTable(container *gabs.Container) {
	msgs, err := container.Search("topMessages").Children()
	if err != nil {
		log.WithField("err", err).Fatal("no messages found.")
		return
	}
	var items [][]string
	for _, msg := range msgs {
		key := getString(msg, "key")
		value := getFloat64(msg, "value")
		items = append(items, []string{
			key,
			fmt.Sprintf("%d", int64(value)),
		})
	}
	u.Table.Rows = items
	ui.Render(u.Table)
}

func (u *UI) updateBars(container *gabs.Container) {
	levels, err := container.Search("levels").ChildrenMap()
	if err != nil {
		return
	}
	levelLen := len(levels)
	data := make([]int, levelLen)
	labels := make([]string, levelLen)

	for k, val := range levels {
		count := val.Data().(float64)
		data = append(data, int(count))
		labels = append(labels, k)
	}

	u.Bars.Data = data
	u.Bars.DataLabels = labels
	ui.Render(u.Bars)
}

func getString(c *gabs.Container, path string) string {
	if v, ok := c.Search(path).Data().(string); ok {
		return v
	}
	return ""
}

func getFloat64(c *gabs.Container, path string) float64 {
	if v, ok := c.Search(path).Data().(float64); ok {
		return v
	}
	return 0
}
