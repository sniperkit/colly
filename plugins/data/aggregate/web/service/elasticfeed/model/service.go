package model

import (
	"github.com/sniperkit/colly/plugins/data/aggregate/service/service/stream"
)

type ServiceManager interface {
	GetStreamService() *stream.StreamService

	Init()
}
