package system

import (
	"github.com/sniperkit/colly/plugins/data/aggregate/service/service/system/v1/router"
)

type SystemService struct{}

func (this *SystemService) Init() {
	router.InitRouters()
	router.InitPluginRouters()
}

func NewSystemService() *SystemService {
	return &SystemService{}
}
