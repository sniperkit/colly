package runtime

import (
	"fmt"
	goruntime "runtime"
	"strings"
	"github.com/v2pro/plz/countlog"
)

type JavascriptError struct {
	funcName    string
	targetJsSrc string
	targetGoSrc string
	Recovered   interface{}
}

func (err *JavascriptError) Error() string {
	return err.String()
}

func (err *JavascriptError) String() string {
	return fmt.Sprintf("function:\n%s\njavascript source:\n%s\n===\ngo source:\n%s\n===\nerror: %v",
		err.funcName, err.targetJsSrc, err.targetGoSrc, err.Recovered)
}

func (err *JavascriptError) MarshalJSON() ([]byte, error) {
	stream := Json.BorrowStream(nil)
	defer Json.ReturnStream(stream)
	stream.WriteObjectStart()
	stream.WriteObjectField("function")
	stream.WriteVal(err.funcName)
	stream.WriteMore()
	stream.WriteObjectField("javascript source")
	stream.WriteVal(err.targetJsSrc)
	stream.WriteMore()
	stream.WriteObjectField("go source")
	stream.WriteVal(err.targetGoSrc)
	stream.WriteMore()
	stream.WriteObjectField("error")
	stream.WriteVal(err.Recovered)
	stream.WriteObjectEnd()
	return stream.Buffer(), stream.Error
}

type Thrown struct {
	Value interface{}
}

func (thrown Thrown) ErrorNumber() int {
	return 20002
}

func (thrown Thrown) Error() string {
	return fmt.Sprintf("%v", thrown.Value)
}

func (thrown Thrown) MarshalJSON() ([]byte, error) {
	return Json.Marshal(thrown.Value)
}

func ReportError(funcName string, jsSrc string, jsOffsets []int, goSrc string, goOffsets []int,
	absoluteLineNo int, recovered interface{}) {
	if _, isJsError := recovered.(*JavascriptError); isJsError {
		panic(recovered)
	}
	targetFile, frames := listFrames()
	if frames == nil {
		panic(&JavascriptError{
			funcName:  funcName,
			Recovered: recovered,
		})
	}
	line := searchFrame(frames, targetFile, funcName)
	if line == -1 {
		panic(&JavascriptError{
			funcName:  funcName,
			Recovered: recovered,
		})
	}
	targetLine := line - absoluteLineNo - 9
	targetGoSrc, targetStart, targetEnd := searchLine(goSrc, targetLine)
	if targetGoSrc == "" {
		panic(&JavascriptError{
			funcName:  funcName,
			Recovered: recovered,
		})
	}
	startIndex, endIndex := searchGoOffsets(goOffsets, targetStart, targetEnd)
	jsStart := jsOffsets[startIndex]
	var targetJsSrc string
	if endIndex == -1 {
		targetJsSrc = jsSrc[jsStart:]
	} else {
		jsEnd := jsOffsets[endIndex]
		targetJsSrc = jsSrc[jsStart:jsEnd]
	}
	targetJsSrc = strings.TrimSpace(targetJsSrc)
	panic(&JavascriptError{
		funcName:    funcName,
		targetGoSrc: targetGoSrc,
		targetJsSrc: targetJsSrc,
		Recovered:   recovered,
	})
}

func listFrames() (string, []*goruntime.Frame) {
	pc := make([]uintptr, 256)
	pc = pc[:goruntime.Callers(2, pc)]
	framesIter := goruntime.CallersFrames(pc)
	suffix := ".Fn" // entry point of generated code
	frames := []*goruntime.Frame{}
	for {
		frame, more := framesIter.Next()
		if len(frame.Function) >= len(suffix) &&
			frame.Function[len(frame.Function)-len(suffix):] == suffix {
			return frame.File, frames
		}
		frames = append(frames, &frame)
		if !more {
			break
		}
	}
	return "", nil
}

func searchFrame(frames []*goruntime.Frame, targetFile string, funcName string) int {
	suffix := "." + funcName
	for _, frame := range frames {
		if frame.File != targetFile {
			continue
		}
		if strings.HasSuffix(frame.Function, suffix) {
			countlog.Trace("event!runtime.found panic frame", "funcName", funcName,
				"file", frame.File, "line", frame.Line)
			return frame.Line
		}
	}
	countlog.Warn("event!runtime.panic frame not found", "funcName", funcName)
	return -1
}

func searchLine(str string, targetLine int) (string, int, int) {
	tmp := str
	var offset int
	var start int
	for i := 0; i < targetLine; i++ {
		start = strings.IndexByte(tmp, '\n')
		if start == -1 {
			return "", 0, 0
		}
		tmp = tmp[start+1:]
		offset += start + 1
	}
	end := strings.IndexByte(tmp, '\n')
	if end == -1 {
		end = len(tmp)
	}
	return tmp[:end], offset, offset + end
}

func searchGoOffsets(goOffsets []int, targetStart int, targetEnd int) (int, int) {
	var startFound bool
	var startIndex int
	for i := 0; i < len(goOffsets)-1; i++ {
		if targetStart >= goOffsets[i] && targetStart < goOffsets[i+1] {
			startFound = true
			startIndex = i
		}
		if startFound && goOffsets[i+1] >= targetEnd {
			return startIndex, i + 1
		}
	}
	if startFound {
		return startIndex, -1
	}
	return len(goOffsets) - 1, -1
}
