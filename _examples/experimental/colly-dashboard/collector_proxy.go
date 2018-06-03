package main

import (
	proxy "github.com/sniperkit/colly/pkg/proxy/default"
)

// collector - proxy
var (
	cp  colly.ProxyFunc  // collector's default proxy function
	opl *onion.ProxyList // collector's multi-protocol proxy object
	// opf onion.ProxyFunc // collector's multi-protocol proxy function
)
