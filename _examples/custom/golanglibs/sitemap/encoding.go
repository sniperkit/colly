package main

import (
	// "bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/pquerna/ffjson/ffjson"

	// "https://github.com/chonla/dbz/blob/master/db/sqlite.go"
	// "github.com/cnf/structhash"
	// "github.com/siddontang/go-mysql-elasticsearch"
	// "github.com/mandolyte/csv-utils"
	// pp "github.com/sniperkit/xutil/plugin/debug/pp"

	"github.com/sniperkit/xutil/plugin/format/convert/json2csv"
	jsoniter "github.com/sniperkit/xutil/plugin/format/json"
)

/*
	Refs:
	- https://github.com/dfontana/GaggleOfKaggle/blob/master/Scraping/go_scrap.go
	- https://github.com/Gujarats/csv-reader/blob/master/app.go
	- https://golang.org/pkg/encoding/csv/#example_Reader
	- https://github.com/janlay/text-builder/blob/master/text-builder.go
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

// csv reader/writer defaults
var (
	csvSplitAt             int      = 2500
	csvReaderBuffer        int      = 20000
	csvStreamBuffer        int      = 20000
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

type csvStream struct {
	path         string `required:'true'` // Filepath to Local CSV File
	splitAt      int    `default:'2500'`
	buffer       int    `default:'20000'`
	lock         *sync.Mutex
	wg           *sync.WaitGroup
	selectorType string   `default:'name'` // column selector type, availble: by_key, by_name (default: column_name)
	columnsKeys  []int    // default: 0
	columnsNames []string // default: "url"
	paused       bool
	ready        bool
	debug        bool
	isRemote     bool
	// file         *os.File
	reader *csv.Reader
	*csvLine
}

func NewStreamCSV(fp string, st string) (*csvStream, error) {

	s := &csvStream{
		isRemote: isRemoteURL(fp),
		path:     fp,
		buffer:   20000,
		splitAt:  2500,
		lock:     &sync.Mutex{},
		wg:       &sync.WaitGroup{},
	}

	st = strings.ToLower(st)
	switch st {
	case "by_key":
		s.selectorType = st
		s.columnsKeys = append(s.columnsKeys, 0)
	case "by_name":
		fallthrough
	default:
		s.selectorType = st
		s.columnsNames = append(s.columnsNames, "url")
	}

	// var reader *csv.Reader
	if !s.isRemote {
		if _, err := os.Stat(s.path); os.IsNotExist(err) {
			return nil, err
		}
		log.Infoln("reading file:", s.path)
		file, err := os.Open(s.path)
		if err != nil {
			log.Fatalln("failed to open file, error: ", err)
			return nil, err
		}
		defer file.Close()
		s.reader = csv.NewReader(file)

	} else {
		log.Infoln("loading remote:", s.path)
		resp, err := http.Get(s.path)
		if err != nil {
			log.Fatalln("failed to fetch content, error: ", err)
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			log.Fatalln("failed with status", resp.Status)
			return nil, errInvalidRemoteStatusCode
		}
		s.reader = csv.NewReader(resp.Body)

	}

	// s.reader.Comma = ';'
	// pp.Println("StreamCSV=", s)

	s.ready = true
	return s, nil
}

func (cs *csvStream) Wait() {
	// task.lock.Lock()
	// defer task.lock.Unlock()
	cs.wg.Wait()
}

// func (cs *csvStream) Read(csv *csv.Reader, buffer int) (lines chan *csvLine) {
func (cs *csvStream) ReadAsync() (lines chan *csvLine) {
	// cs.lock.Lock()
	// defer cs.lock.Unlock()

	lines = make(chan *csvLine, cs.buffer)
	// cs.wg.Add(1)
	go func() {
		header, err := cs.reader.Read()
		if err != nil {
			close(lines)
			return
		}
		i := 0
		for {
			line, err := cs.reader.Read()
			if len(line) > 0 {
				i++
				lines <- &csvLine{
					header: header,
					line:   line,
				}
				log.Println("header=", header, ", line=", line)
			}
			if err == io.EOF {
				// cs.wg.Done()
				return
			}
			if err != nil {
				log.Printf("Sent %d lines\n", i)
				close(lines)
				// cs.wg.Done()
				return
			}
		}
	}()
	cs.wg.Wait()
	return
}

/*
func (cs *csvStream) Scan() (lines chan *csvLine) {
	// cs.lock.Lock()
	// defer cs.lock.Unlock()

	var reader io.Reader
	if !cs.isRemote {
		if _, err := os.Stat(cs.path); os.IsNotExist(err) {
			return
		}
		log.Infoln("reading file:", cs.path)
		fi, err := os.Open(cs.path)
		if err != nil {
			log.Fatalln("failed to open file, error: ", err)
			return
		}
		defer fi.Close()
		tabular = csv.NewReader(fi)

	} else {
		log.Infoln("loading remote:", cs.path)
		resp, err := http.Get(cs.path)
		if err != nil {
			log.Fatalln("failed to fetch content, error: ", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			log.Fatalln("failed with status", resp.Status)
			return
		}
		reader = csv.NewReader(resp.Body)

	}

	// csv.Comma = ';'
	pp.Println("StreamCSV=", cs)
	lines = make(chan *csvLine, cs.buffer)
	scanner := bufio.NewScanner(reader)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) > 0 {
				i++
				lines <- &csvLine{
					header: header,
					line:   line,
				}
				log.Println("header=", header, ", line=", line)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatalln("csv scanner error=", err)
			return
		}

	}()

	return
}
*/

func (cs *csvStream) Pause(status bool) {
	cs.lock.Lock()
	cs.paused = status
	cs.lock.Unlock()
}

/*
func (cs *csvStream) Close() {
	cs.lock.Lock()
	cs.file.Close()
	cs.lock.Unlock()
}
*/

func (cs *csvStream) SplitAt(limit int) *csvStream {
	cs.lock.Lock()
	if limit <= 0 {
		limit = csvSplitAt
	}
	cs.splitAt = limit
	cs.lock.Unlock()
	return cs
}

func (cs *csvStream) Buffer(buffer int) *csvStream {
	cs.lock.Lock()
	cs.buffer = buffer
	cs.lock.Unlock()
	return cs
}

func (cs *csvStream) SetColumnsKeys(keys ...int) *csvStream {
	cs.lock.Lock()
	cs.columnsKeys = keys
	cs.lock.Unlock()
	return cs
}

func (cs *csvStream) SetColumnsNames(names ...string) *csvStream {
	cs.lock.Lock()
	cs.columnsNames = names
	cs.lock.Unlock()
	return cs
}

/*
func (cl *csvLine) Get(idx int) (value string) {
	return strings.TrimSpace(cl.line[idx])
}

func (cl *csvLine) GetByName(key string) (value string) {
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
*/

//func (cs *csvStream) hasHeader() *csvStream {
//	return cs
//}

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

func (cl *csvLine) GetByKey(k int) (value string) {
	return strings.TrimSpace(cl.line[k])
}

func (cl *csvLine) GetByName(key string) (value string) {
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
	cft.created_at = csv.GetByName("created_at")
	cft.duration = csv.GetByName("duration")
	cft.duration_time = csv.GetByName("duration_time")
	cft.finished_at = csv.GetByName("finished_at")
	cft.service = csv.GetByName("service")
	cft.topic = csv.GetByName("topic")
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
