package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/sniperkit/xstats/pkg"
	"github.com/sniperkit/xtask/plugin/counter"
	// pp "github.com/sniperkit/xutil/plugin/debug/pp"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	logTasks           bool = true
	FullyQualifiedPath bool = false
)

var (
	log                         = logrus.New()
	counters     *counter.Oc    = counter.NewOc()
	counterAsync map[string]int = make(map[string]int)
)

func init() {
	log.Formatter = new(prefixed.TextFormatter)
	log.Level = logrus.DebugLevel
}

// type logFields logrus.Fields
type Fields logrus.Fields

// WithFields is an alias for logrus.WithFields.
func LogWithFields(f Fields) *logrus.Entry {
	return logrus.WithFields(logrus.Fields(f))
}

type funcMetrics struct {
	calls struct {
		count  int           `metric:"count" type:"counter"`
		failed int           `metric:"failed" type:"counter"`
		time   time.Duration `metric:"time"  type:"histogram"`
	} `metric:"func.calls"`
}

func GetCaller() string {
	_, file, line, _ := runtime.Caller(3)
	return fmt.Sprintf("%s:%d", trimPath(file), line)
}

func timeTrack(startedAt time.Time, topic string) { //, reqInfo map[string]interface{}
	c := GetCaller()
	go func(c string) {

		elapsed := time.Since(startedAt)

		if logTasks {
			completedAt := startedAt.Add(elapsed)
			expiredAt := completedAt.Add(cacheTTL)

			task := make(map[string]interface{}, 10)
			task["topic"] = topic
			task["tags"] = []string{}
			task["service"] = "github"
			task["task_duration"] = elapsed.Seconds()
			task["task_creation_datetime"] = startedAt.String()
			task["task_creation_timestamp"] = strconv.FormatInt(startedAt.UTC().Unix(), 10)
			task["task_completed_datetime"] = completedAt.String()
			task["task_completed_timestamp"] = strconv.FormatInt(completedAt.UTC().Unix(), 10)
			task["task_expired_datetime"] = expiredAt
			task["task_expired_timestamp"] = strconv.FormatInt(expiredAt.UTC().Unix(), 10)
			cds.Append("tasks", task)
		}

		log.Printf("main().timeTrack() %s took %s / %s", topic, elapsed, c)
	}(c)
}

// See http://stackoverflow.com/a/7053871/199475
func Function(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func trimPath(path string) string {
	// For details, see https://github.com/uber-go/zap/blob/e15639dab1b6ca5a651fe7ebfd8d682683b7d6a8/zapcore/entry.go#L101
	if idx := strings.LastIndexByte(path, '/'); idx >= 0 {
		if idx := strings.LastIndexByte(path[:idx], '/'); idx >= 0 {
			// Keep everything after the penultimate separator.
			return path[idx+1:]
		}
	}
	return path
}

func addMetrics(start time.Time, incr int, failed bool) {
	callTime := time.Now().Sub(start)
	m := &funcMetrics{}
	m.calls.count = incr
	m.calls.time = callTime
	if failed {
		m.calls.failed = incr
	}
	stats.Report(m)
}

func funcTrack(start time.Time) {
	return
	function, file, line, _ := runtime.Caller(1)
	go func() {
		elapsed := time.Since(start)
		log.Printf("main().funcTrack() %s took %s", fmt.Sprintf("%s:%s:%d", runtime.FuncForPC(function).Name(), chopPath(file), line), elapsed)
	}()
}

func counterTrack(name string, incr int) {
	go func() {
		counters.Increment(name, incr)
	}()
}
