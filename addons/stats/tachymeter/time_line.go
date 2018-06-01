package tachymeter

import (
	"strconv"
	"strings"
	"time"
)

// Timeline holds a []*timelineEvents,
// which nest *Metrics for analyzing
// multiple collections of measured events.
type Timeline struct {
	timeline []*timelineEvent
}

// AddEvent adds a *Metrics to the *Timeline.
func (t *Timeline) AddEvent(m *Metrics) {
	t.timeline = append(t.timeline, &timelineEvent{
		Metrics: m,
		Created: time.Now(),
	})
}

// timelineEvent holds a *Metrics and
// time that it was added to the Timeline.
type timelineEvent struct {
	Metrics *Metrics
	Created time.Time
}

// genGraphHTML takes a *timelineEvent and id (used for each graph
// html element ID) and creates a chart.js graph output.
func genGraphHTML(te *timelineEvent, id int) string {
	keys := []string{}
	values := []uint64{}

	for _, b := range *te.Metrics.Histogram {
		for k, v := range b {
			keys = append(keys, k)
			values = append(values, v)
		}
	}

	keysj, _ := json.Marshal(keys)
	valuesj, _ := json.Marshal(values)

	out := strings.Replace(graph, "XCANVASID", strconv.Itoa(id), 1)
	out = strings.Replace(out, "XKEYS", string(keysj), 1)
	out = strings.Replace(out, "XVALUES", string(valuesj), 1)

	return out
}
