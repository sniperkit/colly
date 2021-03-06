package onion

import (
	"fmt"
	"sync"
)

type ProxyPool interface {
	SetPool() error
}

type Proxy struct {
	p     string `required:'true'`
	ready bool   `default:'false'`
	isTor bool   `default:'false'`
	lock  *sync.RWMutex
	err   error
}

// Constructor
func New(ps string) *Proxy {
	return &Proxy{
		p:    ps,
		lock: &sync.RWMutex{},
	}
}

// Constructor
func NewWithConfig(c *Config) *Proxy {
	ps := fmt.Sprintf("%s://%s:%d", c.Protocol, c.Host, c.Port)
	p := &Proxy{
		p:    ps,
		lock: &sync.RWMutex{},
	}
	if c.Port == 9050 {
		p.isTor = true
	}
	// To do:
	// - check if healthy
	// - check if valid
	// - fetch proxy list...
	return p
}

func (p *Proxy) IsReady() bool {
	return bool(p.ready)
}

func (p *Proxy) IsOnion() bool {
	return bool(p.isTor)
}

func (p *Proxy) IsHealthy() (ok bool) {
	if p.err == nil {
		ok = true
	}
	return
}

func (p *Proxy) String() string {
	return string(p.p)
}
