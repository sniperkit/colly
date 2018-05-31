package requests

import (
	"errors"
)

var (
	errCurrentProxyUnset = errors.New("currentProxy is unset")
	errEmptyPoolProxies  = errors.New("Empty pool of proxies")
)
