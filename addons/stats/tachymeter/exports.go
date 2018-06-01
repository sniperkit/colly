package tachymeter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	tabular "github.com/sniperkit/colly/addons/convert/agnostic-tablib"
	fsutil "github.com/sniperkit/xtask/util/fs"
	jsoniter "github.com/sniperkit/xutil/plugin/format/json"
)

var (
	json     = jsoniter.ConfigCompatibleWithStandardLibrary
	dataSet  *tabular.Dataset
	dataBook *tabular.Databook
)

/*
	Refs:
	- https://github.com/agrison/go-tablib
	- github.com/sniperkit/colly/addons/convert/agnostic-tablib
	- github.com/sniperkit/xtask/plugin
*/

type Export struct {
	Encoding   string `default:'tsv'`
	PrefixPath string `default:'./shared/exports/stats/tachymeter/'`
	Basename   string `default:'tachymeter_export'`
	SplitLimit int    `default:'2500'`
	BufferSize int    `default:'20000'`
	BackupMode bool   `default:'true'`
	Overwrite  bool   `default:'true'`
	buffer     *bytes.Buffer
}

// WriteHTML takes an absolute path p and writes an
// html file to 'p/tachymeter-<timestamp>.html' of all
// histograms held by the *Timeline, in series.
func (t *Timeline) WriteHTML(p string) error {
	fsutil.EnsureDir(p)
	path, err := filepath.Abs(p)
	if err != nil {
		return err
	}
	var b bytes.Buffer

	b.WriteString(head)

	// Append graph + info entry for each timeline
	// event.
	for n := range t.timeline {
		// Graph div.
		b.WriteString(fmt.Sprintf(`%s<div class="graph">%s`, tab, nl))
		b.WriteString(fmt.Sprintf(`%s%s<canvas id="canvas-%d"></canvas>%s`, tab, tab, n, nl))
		b.WriteString(fmt.Sprintf(`%s</div>%s`, tab, nl))
		// Info div.
		b.WriteString(fmt.Sprintf(`%s<div class="info">%s`, tab, nl))
		b.WriteString(fmt.Sprintf(`%s<p><h2>Iteration %d</h2>%s`, tab, n+1, nl))
		b.WriteString(t.timeline[n].Metrics.String())
		b.WriteString(fmt.Sprintf("%s%s</p></div>%s", nl, tab, nl))
	}

	// Write graphs.
	for id, m := range t.timeline {
		s := genGraphHTML(m, id)
		b.WriteString(s)
	}

	b.WriteString(tail)

	// Write file.
	d := []byte(b.String())
	fname := fmt.Sprintf("%s/tachymeter-%d.html", path, time.Now().Unix())
	err = ioutil.WriteFile(fname, d, 0644)
	if err != nil {
		return err
	}

	return nil
}
