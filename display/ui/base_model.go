package ui

import (
	"github.com/gdamore/tcell"
)

type BaseModel struct {
	CursorX       int
	CursorY       int
	EndX          int
	EndY          int
	Width         int
	Height        int
	CursorEnabled bool
	CursorVisible bool
}

// to be overwritten
func (b *BaseModel) GetCell(x, y int) (rune, tcell.Style, []rune, int) {
	return ' ', tcell.StyleDefault, nil, 1
}

func (b *BaseModel) SetCursor(x, y int) {
	b.CursorX = x
	b.CursorY = y
	b.limitCursor()
}

func (b *BaseModel) GetCursor() (int, int, bool, bool) {
	return b.CursorX, b.CursorY, b.CursorEnabled, b.CursorVisible
}

func (b *BaseModel) MoveCursor(offx, offy int) {
	b.CursorX += offx
	b.CursorY += offy
	b.limitCursor()
}

func (b *BaseModel) limitCursor() {
	if b.CursorX > b.Width-1 {
		b.CursorX = b.Width - 1
	}
	if b.CursorY > b.Height-1 {
		b.CursorY = b.Height - 1
	}
	if b.CursorX < 0 {
		b.CursorX = 0
	}
	if b.CursorY < 0 {
		b.CursorY = 0
	}
}

func (b *BaseModel) GetBounds() (int, int) {
	return b.Width, b.Height
}
