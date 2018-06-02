package proxy

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/proxy"

	gcb "github.com/devopsfaith/krakend-circuitbreaker/gobreaker"
)

func TestNewMiddleware_multipleNext(t *testing.T) {
	defer func() {
		if r := recover(); r != proxy.ErrTooManyProxies {
			t.Errorf("The code did not panic\n")
		}
	}()
	NewMiddleware(&config.Backend{})(proxy.NoopProxy, proxy.NoopProxy)
}

func TestNewMiddleware_zeroConfig(t *testing.T) {
	for _, cfg := range []*config.Backend{
		&config.Backend{},
		&config.Backend{ExtraConfig: map[string]interface{}{gcb.Namespace: 42}},
	} {
		resp := proxy.Response{}
		mdw := NewMiddleware(cfg)
		p := mdw(dummyProxy(&resp, nil))

		request := proxy.Request{
			Path: "/tupu",
		}

		for i := 0; i < 100; i++ {
			r, err := p(context.Background(), &request)
			if err != nil {
				t.Error(err.Error())
				return
			}
			if &resp != r {
				t.Fail()
			}
		}
	}
}

func TestNewMiddleware_ok(t *testing.T) {
	resp := proxy.Response{}
	mdw := NewMiddleware(&config.Backend{
		ExtraConfig: map[string]interface{}{
			gcb.Namespace: map[string]interface{}{
				"interval":  100.0,
				"timeout":   100.0,
				"maxErrors": 1.0,
			},
		},
	})
	p := mdw(dummyProxy(&resp, nil))

	request := proxy.Request{
		Path: "/tupu",
	}

	for i := 0; i < 100; i++ {
		r, err := p(context.Background(), &request)
		if err != nil {
			t.Error(err.Error())
			return
		}
		if &resp != r {
			t.Fail()
		}
	}
}

func TestNewMiddleware_ko(t *testing.T) {
	expectedErr := fmt.Errorf("Some error")
	calls := uint64(0)
	mdw := NewMiddleware(&config.Backend{
		ExtraConfig: map[string]interface{}{
			gcb.Namespace: map[string]interface{}{
				"interval":  100.0,
				"timeout":   100.0,
				"maxErrors": 1.0,
			},
		},
	})
	p := mdw(func(_ context.Context, _ *proxy.Request) (*proxy.Response, error) {
		total := atomic.AddUint64(&calls, 1)
		if total > 2 {
			t.Error("This proxy shouldn't been executed!")
		}
		return nil, expectedErr
	})

	request := proxy.Request{
		Path: "/tupu",
	}

	for i := 0; i < 2; i++ {
		r, err := p(context.Background(), &request)
		if err != expectedErr {
			t.Error("error expected")
		}
		if nil != r {
			t.Error("not nil response")
		}
	}
	r, err := p(context.Background(), &request)
	if err == nil || err.Error() != "circuit breaker is open" {
		t.Error("error expected")
	}
	if nil != r {
		t.Error("not nil response")
	}
}