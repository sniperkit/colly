package service

import (
	"github.com/sniperkit/colly/plugins/data/aggregate/core/memcache"
)

func NewCache() *memcache.MemcacheClient {
	return memcache.NewMemcacheClient()
}
