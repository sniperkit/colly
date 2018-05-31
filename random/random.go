package random

import (
	"math/rand"
	"time"

	"github.com/dgryski/go-discreterand"
	"github.com/ernsheong/grand"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomChoices(num, stringLen int) []Choice {
	choices := make([]Choice, num)
	var total float64
	for i := 0; i < num; i++ {
		choices[i] = Choice{
			Value:       grand.GenerateRandomString(stringLen),
			Probability: rand.Float64(),
			Level:       INFO_LEVEL,
		}
		total += choices[i].Probability
	}
	var weightedTotal float64
	for i := 0; i < num; i++ {
		if i == num-1 {
			choices[i].Probability = 1 - weightedTotal
		} else {
			choices[i].Probability = choices[i].Probability / total
		}
		weightedTotal += choices[i].Probability
	}
	return choices
}

type Level int32

const (
	INFO_LEVEL Level = iota
	DEBUG_LEVEL
	WARN_LEVEL
	ERROR_LEVEL
)

type LogItem struct {
	Message string
	Fields  map[string]string
	Level   Level
}

type LogGeneration struct {
	LogItems  *WeightedChoices
	NumFields int
	Keys      *WeightedChoices
	Values    *WeightedChoices
}

func (l *LogGeneration) Generate() LogItem {
	fields := make(map[string]string)
	for i := 0; i < l.NumFields; i++ {
		key := l.Keys.Get()
		value := l.Values.Get()
		fields[key] = value
	}
	msg, lvl := l.LogItems.GetWithLevel()
	return LogItem{
		Message: msg,
		Level:   lvl,
		Fields:  fields,
	}
}

type Choice struct {
	Value       string
	Probability float64
	Level       Level
}

type WeightedChoices struct {
	Values        []string
	Probabilities []float64
	Levels        []Level
	alias         discreterand.AliasTable
}

func NewWeightedChoices(choices []Choice) *WeightedChoices {
	choiceLength := len(choices)
	values := make([]string, choiceLength)
	probabilities := make([]float64, choiceLength)
	currLevels := make([]Level, choiceLength)

	for i, choice := range choices {
		values[i] = choice.Value
		probabilities[i] = choice.Probability
		currLevels[i] = choice.Level
	}

	return &WeightedChoices{
		Values:        values,
		Probabilities: probabilities,
		Levels:        currLevels,
		alias:         discreterand.NewAlias(probabilities, rand.NewSource(time.Now().Unix())),
	}
}

func (w *WeightedChoices) Get() string {
	i := w.alias.Next()
	return w.Values[i]
}

func (w *WeightedChoices) GetWithLevel() (string, Level) {
	i := w.alias.Next()
	return w.Values[i], w.Levels[i]
}

type randomMaybeFixed struct {
	RandomLength    int
	FixedValues     []string
	WeightedChoices []Choice
	weighted        *WeightedChoices
	initalized      bool
}

func (r *randomMaybeFixed) init() {
	r.initalized = true
	r.weighted = NewWeightedChoices(r.WeightedChoices)
}

func (r *randomMaybeFixed) Generate() string {
	if !r.initalized {
		r.init()
	}
	fixedLen := len(r.FixedValues)
	if fixedLen == 0 {
		return grand.GenerateRandomString(r.RandomLength)
	} else if len(r.WeightedChoices) > 0 {
		return r.weighted.Get()
	} else {
		return r.FixedValues[rand.Intn(fixedLen)]
	}
}
