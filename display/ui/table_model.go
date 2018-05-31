package ui

import (
	"github.com/gdamore/tcell"
)

type Row struct {
	Text            string
	Children        []*Row
	ChildrenVisible bool
	Level           int
	Style           tcell.Style
}

func (r *Row) AddChild(child *Row) {
	child.Level = r.Level + 1
	r.Children = append(r.Children, child)
}

func (r *Row) GetText(indent string) string {
	var out string
	if len(r.Children) > 0 {
		if r.ChildrenVisible {
			out = "-"
		} else {
			out = "+"
		}
	} else {
		out = " "
	}
	for i := 0; i < r.Level; i++ {
		out += indent
	}
	return out + r.Text
}

type nestedList struct {
	*BaseModel
	Rows     []*Row
	text     []string
	Indent   string
	flatRows []*Row
}

func (n *nestedList) ToggleVisibility() {
	currRow := n.flatRows[n.CursorY]
	if len(currRow.Children) == 0 {
		return
	}
	currRow.ChildrenVisible = !currRow.ChildrenVisible
	n.Update()
}

func (n *nestedList) GetCell(x, y int) (rune, tcell.Style, []rune, int) {
	var ch rune
	if x < 0 || y < 0 || y >= n.Height || x >= n.Width || y >= len(n.text) || x >= len(n.text[y]) {
		return ch, tcell.StyleDefault, nil, 1
	}
	style := n.flatRows[y].Style
	if y == n.CursorY {
		return rune(n.text[y][x]), style.Reverse(true), nil, 1
	} else {
		return rune(n.text[y][x]), style, nil, 1
	}
}

func (n *nestedList) SetRows(rows []*Row) {
	n.Rows = rows
	n.Update()
}

func (n *nestedList) Update() {
	n.text = make([]string, 0)
	n.flatRows = make([]*Row, 0)
	for _, r := range n.Rows {
		n.processRow(r)
	}
	n.Height = len(n.text)
}

func (n *nestedList) processRow(r *Row) {
	currText := r.GetText("  ")
	if currLen := len(currText); currLen > n.Width {
		n.Width = currLen
	}

	n.text = append(n.text, currText)
	n.flatRows = append(n.flatRows, r)

	if r.ChildrenVisible {
		for _, child := range r.Children {
			n.processRow(child)
		}
	}
}
