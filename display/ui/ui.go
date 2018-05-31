package ui

import (
	"context"
	"fmt"
	"net"

	"time"

	"github.com/Jeffail/gabs"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	log "github.com/sirupsen/logrus"
)

type (
	UI struct {
		app        *views.Application
		window     *window
		rowManager *LogRowManager
	}

	window struct {
		views.Panel
		app *views.Application

		title      *views.TextBar
		cell       *views.CellView
		nestedList *nestedList
	}
)

func NewUI() *UI {
	app := &views.Application{}
	title := views.NewTextBar()
	title.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorTeal).
		Foreground(tcell.ColorWhite))

	title.SetCenter("Top Logs By Level", tcell.StyleDefault)

	currWindow := &window{
		app:   app,
		title: title,
		cell:  views.NewCellView(),
		nestedList: &nestedList{
			BaseModel: &BaseModel{
				CursorEnabled: true,
				CursorVisible: true,
			},
		},
	}
	rows := []*Row{
		&Row{
			Text:            "Attempting to Connect..",
			ChildrenVisible: true,
		},
	}

	currWindow.nestedList.SetRows(rows)
	currWindow.cell.SetModel(currWindow.nestedList)

	currWindow.SetTitle(title)
	currWindow.SetContent(currWindow.cell)
	app.SetRootWidget(currWindow)

	return &UI{
		app:        app,
		window:     currWindow,
		rowManager: NewLogRowManager(),
	}
}

func (w *window) HandleEvent(ev tcell.Event) bool {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyRune:
			switch ev.Rune() {
			case 'Q', 'q':
				w.app.Quit()
				return true
			}
		case tcell.KeyEnter:
			w.nestedList.ToggleVisibility()
		}
	}
	return w.Panel.HandleEvent(ev)
}

func (u *UI) Run() {
	if err := u.app.Run(); err != nil {
		fmt.Println("Failed to run app.")
	}
}

func (u *UI) Init() {
	u.Connect()
}

func (u *UI) Connect() {
	conn := u.getConnection()

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

func (u *UI) getConnection() net.Conn {
	for {
		t := time.Now().Add(time.Second)
		ctx, cancel := context.WithDeadline(context.Background(), t)
		conn, _, err := ws.Dial(ctx, "ws://localhost:6002/ws", nil)
		cancel()

		if err != nil {
			time.Sleep(time.Second)
		} else {
			return conn
		}
	}
}

func (u *UI) update(container *gabs.Container) {
	counts := container.Search("levels")

	children, err := container.Search("logs").ChildrenMap()
	if err != nil {
		return
	}
	for logLevel, child := range children {
		count := int(getFloat64(counts, logLevel))
		u.rowManager.SetCount(logLevel, count)
		pairs, err := child.Children()
		if err != nil {
			return
		}
		u.rowManager.ResetChildren(logLevel)

		for _, pair := range pairs {
			key := getString(pair, "key")
			value := int(getFloat64(pair, "value"))
			u.rowManager.AddChild(logLevel, &Row{
				Text: fmt.Sprintf("%d - %s", value, key),
			})
		}
	}
	u.window.nestedList.SetRows(u.rowManager.GetRows())
	u.app.Refresh()
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
