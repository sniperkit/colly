package main

import (
	"github.com/sniperkit/colly/plugins/data/aggregate/service/elasticfeed"
)

func main() {
	engine := elasticfeed.NewElasticfeed()
	engine.Run()
}
