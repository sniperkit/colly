package main

import ()

// ui - dashboard params
var (
	xTermUI   *cui.TermUI
	xResults  chan tui.WorkResult
	stopTheUI chan bool
)
