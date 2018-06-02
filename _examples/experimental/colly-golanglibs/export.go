package main

import (
	"fmt"
	"strings"
	"time"

	// Data Layer Abstraction
	"github.com/ghetzel/pivot"
	"github.com/ghetzel/pivot/dal"
	"github.com/ghetzel/pivot/mapper"

	// filters - probalistic data-structure
	cuckoo "github.com/seiflotfy/cuckoofilter"
	"github.com/willf/bloom"

	// key-value stores
	"github.com/asdine/storm"
	//"gopkg.in/mgo.v2/bson"

	// datastructure helpers
	cmmap "github.com/sniperkit/colly/plugins/data/structure/map/multi"
	tablib "github.com/sniperkit/colly/plugins/data/transform/tabular"
)

const (
	databaseName = "stats.db"
)

var (
	conn     *storm.DB
	Widgets  mapper.Mapper
	sheets   map[string][]interface{}   = make(map[string][]interface{}, 0)
	datasets map[string]*tablib.Dataset = make(map[string]*tablib.Dataset, 0) // := NewDataset([]string{"firstName", "lastName"})
	cds                                 = cmmap.NewConcurrentMultiMap()
)

// cache lists - bloom filters
var (
	bloomFilterSize uint = 20000 // default: 500000
	bloomFilterKeys uint = 5
	blmflt               = bloom.New(bloomFilterSize, bloomFilterKeys)
)

// cache lists - cuckoo filters
var (
	cuckFilterCapacity uint = 20000 // default: 1000000
	cuckFilter              = cuckoo.NewCuckooFilter(cuckFilterCapacity)
	cuckflt                 = cuckoo.NewDefaultCuckooFilter()
)

var WidgetsSchema = &dal.Collection{
	Name:                   `widgets`,
	IdentityFieldType:      dal.StringType,
	IdentityFieldFormatter: dal.GenerateUUID,
	Fields: []dal.Field{
		{
			Name:        `type`,
			Description: `The type of widget.`,
			Type:        dal.StringType,
			Validator:   dal.ValidateIsOneOf(`foo`, `bar`, `baz`),
			Required:    true,
		}, {
			Name:        `usage`,
			Description: `Short description on how to use this widget.`,
			Type:        dal.StringType,
		}, {
			Name:        `created_at`,
			Description: `When the widget was created.`,
			Type:        dal.TimeType,
			Formatter:   dal.CurrentTimeIfUnset,
		}, {
			Name:        `updated_at`,
			Description: `Last time the widget was updated.`,
			Type:        dal.TimeType,
			Formatter:   dal.CurrentTime,
		},
	},
}

type Widget struct {
	ID        string    `pivot:"id,identity"`
	Type      string    `pivot:"type"`
	Usage     string    `pivot:"usage"`
	CreatedAt time.Time `pivot:"created_at"`
	UpdatedAt time.Time `pivot:"updated_at"`
}

func pivotWidget() {
	// setup a new backend instance based on the supplied connection string
	if backend, err := pivot.NewDatabase(`sqlite:///./shared/storage/sqlite3/colly-golanglibs.db`); err == nil {

		// initialize the backend (connect to/open it)
		if err := backend.Initialize(); err == nil {

			// register models to this database backend
			Widgets = mapper.NewModel(backend, WidgetsSchema)

			// create the model tables if they don't exist
			if err := Widgets.Migrate(); err != nil {
				fmt.Printf("failed to create widget table: %v\n", err)
				return
			}
			// make a new Widget instance, containing the data we want to see
			// the ID field will be populated after creation with the auto-
			// generated UUID.
			newWidget := Widget{
				Type:  `foo`,
				Usage: `A fooable widget.`,
			}

			// insert a widget (ID will be auto-generated because of dal.GenerateUUID)
			if err := Widgets.Create(&newWidget); err != nil {
				fmt.Printf("failed to insert widget: %v\n", err)
				return
			}

			// retrieve the widget using the ID we just got back
			var gotWidget Widget

			if err := Widgets.Get(newWidget.ID, &gotWidget); err != nil {
				fmt.Printf("failed to retrieve widget: %v\n", err)
				return
			}

			fmt.Printf("Got Widget: %#+v", gotWidget)

			// delete the widget
			if err := Widgets.Delete(newWidget.ID); err != nil {
				fmt.Printf("failed to delete widget: %v\n", err)
				return
			}
		} else {
			fmt.Printf("failed to initialize backend: %v\n", err)
			return
		}
	} else {
		fmt.Printf("failed to create backend: %v\n", err)
		return
	}
}

type data struct {
	ID        int
	CreatedAt int64
	Market    string
	Size      float64
	SizeStr   string
}

type capdata []data

func (c capdata) Len() int              { return len(c) }
func (c capdata) Swap(i, j int)         { c[i], c[j] = c[j], c[i] }
func (c capdata) Less(i, j int) bool    { return c[i].Size > c[j].Size }
func (c capdata) ToCap(i int) string    { return c[i].SizeStr }
func (c capdata) ToMarket(i int) string { return fmt.Sprintf("%s   ", c[i].Market) }

func (c capdata) Total() string {
	var total float64
	for _, v := range c {
		total += v.Size
	}
	return " $" + seperateFloat(total)
}

func seperateFloat(f float64) (currency string) {
	var fstr = fmt.Sprintf("%0.f", f)
	var offset = 3
	for i := len(fstr); i > 0; i -= 3 {
		if i < 3 {
			offset = i
		}
		sliceText := fstr[i-offset : i]
		currency = strings.TrimSuffix(sliceText+","+currency, ",")
	}
	return
}
