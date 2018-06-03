package main

import ()

// stats - tachymeter params
var (
	startedAt            time.Time
	isTachymeter         bool = true
	isTachymeterParallel bool = false
	cTachymeter          chan *ta.Tachymeter
	xTachy               *ta.Tachymeter
	xTachyResults        *ta.Metrics
	xTachyTimeline       ta.Timeline
	wallTimeStart        time.Time
)
