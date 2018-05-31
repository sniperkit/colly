package logcounter

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"github.com/go-logfmt/logfmt"
	"github.com/timjchin/streaming-counter"
)

type LogCounterConfig struct {
	// number of items to store
	NumResults int
	// Accuracy of the counts within a factor of Epsilon.
	Epsilon float64
	// Porability that the counts are within a factor of Epsilon.
	Delta      float64
	PrintInput bool
}

type Modifier interface {
	OnLineCompletion()
	// Get name of entry in map[string]interface{}
	GetName() string
	// Get interface{} value of the modifier, will be serialized as JSON
	GetPayload() interface{}
}
type ModifierFunc func(string, string)
type LogCounter struct {
	config       *LogCounterConfig
	LevelCounts  map[string]int
	decoder      *logfmt.Decoder
	keyMap       map[string][]ModifierFunc
	allModifiers []Modifier
}

func NewLogCounter(config *LogCounterConfig) *LogCounter {
	l := &LogCounter{
		config:       config,
		keyMap:       make(map[string][]ModifierFunc),
		allModifiers: make([]Modifier, 0),
	}
	t := NewTopMessagesModifier(config)
	levelCounter := NewLevelCounterModifier(config)
	l.AddModifier(levelCounter)
	l.AddModifierFunc("level", levelCounter.OnLevel)

	l.AddModifier(t)
	l.AddModifierFunc("level", t.onLevel)
	l.AddModifierFunc("msg", t.onMessage)
	return l
}

func (l *LogCounter) AddModifier(mod Modifier) {
	l.allModifiers = append(l.allModifiers, mod)
}

func (l *LogCounter) AddModifierFunc(key string, modFunc ModifierFunc) {
	l.keyMap[key] = append(l.keyMap[key], modFunc)
}

func (l *LogCounter) GetState() map[string]interface{} {
	out := make(map[string]interface{})
	for _, mod := range l.allModifiers {
		out[mod.GetName()] = mod.GetPayload()
	}
	return out
}

func (l *LogCounter) ParseReader(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	byteReader := bytes.NewReader([]byte(""))

	for scanner.Scan() {
		currLine := scanner.Bytes()
		byteReader.Reset(currLine)

		dec := logfmt.NewDecoder(byteReader)

		for dec.ScanRecord() {
			for dec.ScanKeyval() {
				k := string(dec.Key())
				v := string(dec.Value())

				if l.keyMap[k] != nil {
					for _, mod := range l.keyMap[k] {
						mod(k, v)
					}
				}
			}
		}

		for _, mod := range l.allModifiers {
			mod.OnLineCompletion()
		}

		if l.config.PrintInput {
			fmt.Println(string(currLine))
		}
	}
}

type Level int32

const (
	UNKNOWN_LEVEL Level = iota
	INFO_LEVEL
	DEBUG_LEVEL
	WARN_LEVEL
	ERROR_LEVEL
)

var LevelStringToLevel = map[string]Level{
	"info":  INFO_LEVEL,
	"debug": DEBUG_LEVEL,
	"warn":  WARN_LEVEL,
	"error": ERROR_LEVEL,
}

var LevelToLevelString = map[Level]string{
	UNKNOWN_LEVEL: "unknown",
	INFO_LEVEL:    "info",
	DEBUG_LEVEL:   "debug",
	WARN_LEVEL:    "warn",
	ERROR_LEVEL:   "error",
}

// Counts the top messages, matching "msg"
type TopMessagesModifier struct {
	config        *LogCounterConfig
	levelMessages map[Level]*counter.StreamingCounter
	currMessage   string
	currLevel     Level
}

func NewTopMessagesModifier(config *LogCounterConfig) *TopMessagesModifier {
	streamingConfig := &counter.StreamingCounterConfig{
		NumResults: config.NumResults,
	}
	return &TopMessagesModifier{
		config: config,
		levelMessages: map[Level]*counter.StreamingCounter{
			UNKNOWN_LEVEL: getCounter(streamingConfig),
			DEBUG_LEVEL:   getCounter(streamingConfig),
			INFO_LEVEL:    getCounter(streamingConfig),
			WARN_LEVEL:    getCounter(streamingConfig),
			ERROR_LEVEL:   getCounter(streamingConfig),
		},
	}
}

func getCounter(streamingConfig *counter.StreamingCounterConfig) *counter.StreamingCounter {
	c, _ := counter.NewStreamingCounter(streamingConfig)
	return c
}

func (t *TopMessagesModifier) onMessage(k, v string) {
	t.currMessage = v
}

func (t *TopMessagesModifier) onLevel(k, v string) {
	t.currLevel = LevelStringToLevel[v]
}

func (t *TopMessagesModifier) GetName() string {
	return "logs"
}

func (t *TopMessagesModifier) GetPayload() interface{} {
	outMap := map[string]interface{}{}
	for lvl, count := range t.levelMessages {
		outMap[LevelToLevelString[lvl]] = count.GetAll()
	}
	return outMap
}

func (t *TopMessagesModifier) OnLineCompletion() {
	t.levelMessages[t.currLevel].Add(t.currMessage)
}

// Counts the number of log levels seen.
type LevelCounterModifier struct {
	LevelCounts map[string]int
}

func NewLevelCounterModifier(config *LogCounterConfig) *LevelCounterModifier {
	return &LevelCounterModifier{
		LevelCounts: make(map[string]int),
	}
}

func (t *LevelCounterModifier) OnLevel(k, v string) {
	t.LevelCounts[v]++
}

func (t *LevelCounterModifier) GetName() string {
	return "levels"
}

func (t *LevelCounterModifier) GetPayload() interface{} {
	return t.LevelCounts
}

func (t *LevelCounterModifier) OnLineCompletion() {}
