package ui

import (
	"fmt"

	"github.com/gdamore/tcell"
)

type LogRowManager struct {
	rows     map[string]*Row
	currRows []*Row
}

func NewLogRowManager() *LogRowManager {
	l := &LogRowManager{
		rows: map[string]*Row{
			"info":    &Row{ChildrenVisible: true, Style: tcell.StyleDefault.Foreground(tcell.ColorGreen)},
			"debug":   &Row{ChildrenVisible: true, Style: tcell.StyleDefault.Foreground(tcell.ColorBlue)},
			"warn":    &Row{ChildrenVisible: true, Style: tcell.StyleDefault.Foreground(tcell.ColorYellow)},
			"error":   &Row{ChildrenVisible: true, Style: tcell.StyleDefault.Foreground(tcell.ColorRed)},
			"unknown": &Row{ChildrenVisible: true, Style: tcell.StyleDefault.Foreground(tcell.ColorGreen)},
		},
	}
	l.currRows = []*Row{
		l.rows["error"],
		l.rows["warn"],
		l.rows["info"],
		l.rows["debug"],
		l.rows["unknown"],
	}
	return l
}

func (l *LogRowManager) SetCount(k string, count int) {
	row := l.rows[k]
	if row == nil {
		return
	}
	row.Text = fmt.Sprintf("Level: %s \t Count:%d", k, count)
}

func (l *LogRowManager) ResetChildren(k string) {
	row := l.rows[k]
	if row == nil {
		return
	}
	row.Children = []*Row{}
}

func (l *LogRowManager) AddChild(k string, childRow *Row) {
	row := l.rows[k]
	if row == nil {
		return
	}
	row.AddChild(childRow)
}

func (l *LogRowManager) GetRows() []*Row {
	return l.currRows
}
