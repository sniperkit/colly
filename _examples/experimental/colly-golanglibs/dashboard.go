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

	// core
	colly "github.com/sniperkit/colly/pkg"

	// experimental addons
	//// console UI
	cui "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/gocui"
	tui "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/termui"
	// tui_hist "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/termui/histogram"
	// dash "github.com/sniperkit/colly/plugins/cmd/dashboard"
	// cui "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/gocui"
	// tvi "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/tview"
)

/*
	Refs:
	- https://github.com/c1982/mcap/blob/master/main.go
*/

var (
	xConsoleUI *cui.TermUI
	xTermUI    *cui.TermUI
	xResults   chan tui.WorkResult
	stopTheUI  chan bool
)
