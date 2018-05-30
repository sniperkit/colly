package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/pquerna/ffjson/ffjson"
	// "https://github.com/chonla/dbz/blob/master/db/sqlite.go"
	// "github.com/cnf/structhash"
	// "github.com/siddontang/go-mysql-elasticsearch"
	// "github.com/mandolyte/csv-utils"

	"github.com/sniperkit/xutil/plugin/format/convert/json2csv"
	jsoniter "github.com/sniperkit/xutil/plugin/format/json"
)

/*
	Refs:
	- https://github.com/dfontana/GaggleOfKaggle/blob/master/Scraping/go_scrap.go
	- https://github.com/Gujarats/csv-reader/blob/master/app.go
*/

const (
	DEFAULT_CSV_READER_BUFFER           int = 20000
	DEFAULT_CSV_STREAM_BUFFER           int = 20000
	DEFAULT_CSV_STREAM_OUTPUT_MAX_LINES int = 1000
)

// global export default variables
var (
	defaultExportPrefixPath = cachePrefixPath + "./shared/storage/export"
)

// json reader/writer defaults
var ()

// csv reader/writer defaults
var (
	csvReaderBuffer        int      = 20000
	csvCtreamBuffer        int      = 20000
	csvStreamOuputMaxLines int      = 1000
	csvStreamOuputColumns  []string = []string{"domain", "loc", "created_at", "duration", "duration_time", "finished_at"}
)

type CsvWriter struct {
	csvWriter *csv.Writer
	lock      *sync.Mutex
	wg        *sync.WaitGroup
}

func newSafeCsvWriter(fileName string) (*CsvWriter, error) {
	csvFile, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	w := csv.NewWriter(csvFile)
	return &CsvWriter{
		csvWriter: w,
		lock:      &sync.Mutex{},
		wg:        &sync.WaitGroup{},
	}, nil
}

func (w *CsvWriter) Write(row []string) {
	w.lock.Lock()
	w.csvWriter.Write(row)
	w.lock.Unlock()
}

func (w *CsvWriter) Flush() {
	w.lock.Lock()
	w.csvWriter.Flush()
	w.lock.Unlock()
}

// streamCsv
//  Streams a CSV Reader into a returned channel.  Each CSV row is streamed along with the header.
//  "true" is sent to the `done` channel when the file is finished.

type csvLine struct {
	header []string
	line   []string
}

// Args
//  csv    - The csv.Reader that will be read from.
//  buffer - The "lines" buffer factor.  Send "0" for an unbuffered channel.
func streamCsv(csv *csv.Reader, buffer int) (lines chan *csvLine) {
	lines = make(chan *csvLine, buffer)
	go func() {
		header, err := csv.Read()
		if err != nil {
			close(lines)
			return
		}
		i := 0
		for {
			line, err := csv.Read()
			if len(line) > 0 {
				i++
				lines <- &csvLine{
					header: header,
					line:   line,
				}
			}
			if err != nil {
				log.Printf("Sent %d lines\n", i)
				close(lines)
				return
			}
		}
	}()
	return
}

func (cl *csvLine) Get(key string) (value string) {
	x := -1
	for i, value := range cl.header {
		if value == key {
			x = i
			break
		}
	}
	if x == -1 {
		return ""
	}
	return strings.TrimSpace(cl.line[x])
}

type csvFlowLine struct {
	created_at    string
	duration      string
	duration_time string
	finished_at   string
	service       string
	topic         string
}

// TODO: Find a better way
func streamFormatCsvToLine(csv *csvLine) (*csvFlowLine, error) {
	cft := csvFlowLine{}
	/*
		// How to call struct attribute dynamically ?!
		for _, c := range csvStreamOuputColumns {
			cft.created_at = csv.Get("c")
		}
	*/
	cft.created_at = csv.Get("created_at")
	cft.duration = csv.Get("duration")
	cft.duration_time = csv.Get("duration_time")
	cft.finished_at = csv.Get("finished_at")
	cft.service = csv.Get("service")
	cft.topic = csv.Get("topic")
	return &cft, nil
}

func printCsvFlowTable(lines chan *csvFlowLine) (done chan int) {
	done = make(chan int)
	go func() {
		table := csvFlowTable{}
		i := 0
		for line := range lines {
			i++
			table = append(table, line)
			if len(table) >= csvStreamOuputMaxLines {
				table.Send()
				table = csvFlowTable{}
			}
		}
		if len(table) > 0 {
			table.Send()
		}
		done <- i
	}()
	return
}

func streamCsvLine(csvLines chan *csvLine) (lines chan *csvFlowLine) {
	lines = make(chan *csvFlowLine, csvReaderBuffer)
	go func() {
		var flowLine *csvFlowLine
		for line := range csvLines {
			flowLine, _ = streamFormatCsvToLine(line)
			lines <- flowLine
		}
		close(lines)
	}()
	return
}

type csvFlowTable []*csvFlowLine

func (cft *csvFlowTable) Send() {
	// code to send to the database here.
	log.Printf("----\nSending %d lines\n%s", len(*cft), *cft)
}

// json related encoding helpers
var (
	json                                            = jsoniter.ConfigCompatibleWithStandardLibrary
	writers          map[string]*json2csv.CSVWriter = make(map[string]*json2csv.CSVWriter, 0)
	jsonfile         map[string]*os.File            = make(map[string]*os.File, 0)
	headerStyleTable                                = map[string]json2csv.KeyStyle{
		"jsonpointer": json2csv.JSONPointerStyle,
		"slash":       json2csv.SlashStyle,
		"dot":         json2csv.DotNotationStyle,
		"dot-bracket": json2csv.DotBracketStyle,
	}
)

func getHeaders(filterMap map[string]string) []string {
	var hdrs []string
	for k, _ := range filterMap {
		hdrs = append(hdrs, k)
	}
	return hdrs
}

func initWriters(truncate bool, groups ...string) {
	for _, group := range groups {
		if writers[group] == nil {
			writers[group] = newWriterJSON2CSV(truncate, group)
		}
	}
}

func encode(item interface{}, out io.Writer) {
	buf, err := ffjson.Marshal(&item)
	if err != nil {
		log.Fatalln("Encode error:", err)
	}
	// Write the buffer
	_, _ = out.Write(buf)
	// We are now no longer need the buffer so we pool it.
	ffjson.Pool(buf)
}

func encodeRows(rows []interface{}, out io.Writer) {
	// We create an encoder.
	enc := ffjson.NewEncoder(out)
	for _, item := range rows {
		if item == nil {
			continue
		}
		// Encode into the buffer
		err := enc.Encode(&item)
		enc.SetEscapeHTML(false)
		if err != nil {
			log.Fatalln("encodeItems error:", err)
		}
		// If err is nil, the content is written to out, so we can write to it as well.
		//if i != len(rows)-1 {
		//	_, _ = out.Write([]byte{""})
		//}
	}
}

// https://github.com/fanliao/go-concurrentMap#safely-use-composition-operation-to-update-the-value-from-multiple-threads
/*---- group string by first char using ConcurrentMap ----*/
//sliceAdd function returns a function that appends v into slice
var sliceAdd = func(v interface{}) func(interface{}) interface{} {
	return func(oldVal interface{}) (newVal interface{}) {
		if oldVal == nil {
			vs := make([]string, 0, 1)
			return append(vs, v.(string))
		} else {
			return append(oldVal.([]string), v.(string))
		}
	}
}

func flushWriters() {
	for k, w := range writers {
		if w != nil {
			data, _ := cds.Get(k)

			if len(data) <= 0 {
				continue
			}

			if jsonfile[k] == nil {
				jsonOutpuFile := fmt.Sprintf(defaultExportPrefixPath+"/json/%s.json", k)
				var err error
				jsonfile[k], err = os.OpenFile(jsonOutpuFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Fatalln(" os.Create(jsonOutpuFile) error:", err)
				}
			}
			encodeRows(data, jsonfile[k])
			results, err := json2csv.JSON2CSV(data)
			if err != nil {
				log.Fatalln("JSON2CSV error:", err)
			}
			w.WriteCSV(results)

			w.Flush()
			if err := w.Error(); err != nil {
				log.Fatalln("Error: ", err)
			}
		}
	}
	// jsonfile[k].Close()
	cds.Clear()
}

// add prefixPath, headerStyleTable, transpose
func newWriterJSON2CSV(truncate bool, basename string) *json2csv.CSVWriter {
	outputFile := fmt.Sprintf(defaultExportPrefixPath+"/csv/%s.csv", basename)
	log.Debugln("instanciate new concurrent writer to output file=", outputFile)
	w, err := json2csv.NewCSVWriterToFile(outputFile)
	if err != nil {
		log.Fatalf("Could not open `%s` for writing", outputFile)
	}
	w.HeaderStyle = headerStyleTable["dot"]
	w.NoHeaders(true)
	return w
}
