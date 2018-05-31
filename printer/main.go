package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/timjchin/logcounter/random"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
	})
	log.SetOutput(os.Stdout)
	l := &random.LogGeneration{
		LogItems: random.NewWeightedChoices([]random.Choice{
			{"received a device request", 0.05, random.INFO_LEVEL},
			{"user to access home", 0.5, random.INFO_LEVEL},
			{"user to access about", 0.4, random.INFO_LEVEL},
			{"failed request", 0.05, random.ERROR_LEVEL},
		}),
		NumFields: 1,
		Keys:      random.NewWeightedChoices(random.RandomChoices(50, 10)),
		Values:    random.NewWeightedChoices(random.RandomChoices(500, 10)),
	}
	for {
		item := l.Generate()
		logToLogrus(item)
		time.Sleep(time.Millisecond * time.Duration(200))
	}
}

func logToLogrus(l random.LogItem) {
	f := log.Fields{}
	for k, v := range l.Fields {
		f[k] = v
	}
	currLog := log.WithFields(f)
	switch l.Level {
	case random.INFO_LEVEL:
		currLog.Info(l.Message)
	case random.DEBUG_LEVEL:
		currLog.Debug(l.Message)
	case random.ERROR_LEVEL:
		currLog.Error(l.Message)
	case random.WARN_LEVEL:
		currLog.Warn(l.Message)
	}
}
