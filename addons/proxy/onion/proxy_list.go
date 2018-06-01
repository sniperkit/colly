package onion

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ProxyList struct {
	Pool         []*Proxy
	Source       string
	client       *http.Client
	currentProxy *Proxy
	currentInd   int
	lock         *sync.RWMutex
}

func (pl *ProxyList) Init() error {
	pl.Source = REMOTE_PROXY_LIST_URL
	err := pl.SetPool()
	if err != nil {
		return err
	}
	err = pl.SetProxy()
	if err != nil {
		return err
	}
	return nil
}

func (pl *ProxyList) SetPool() error {
	res, err := OnionRequest(pl.Source)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		panic(err)
	}
	var pool = make([]*Proxy, 10)
	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		elem := s.Find("li").Find("script").Text()
		encoded := strings.Split(elem, "'")
		for _, enc := range encoded {
			pr, err := base64.StdEncoding.DecodeString(enc)
			if err != nil {
				continue
			}
			if pr == nil {
				continue
			}
			ps := string(pr)
			numDots := len(strings.Split(ps, "."))
			if numDots == 4 {
				pool = append(pool, New(ps))
				continue
			}
		}

	})
	pl.Pool = pool
	return nil
}

func (pl *ProxyList) NumProxies() int {
	return len(pl.Pool)
}

func (pl *ProxyList) SetProxy() error {
	n := pl.NumProxies()
	if n == 0 {
		err := errEmptyPoolProxies
		return err
	}
	ind := rand.Intn(n)
	pr := pl.Pool[ind]
	proxyStr := fmt.Sprintf("http://%s", pr)
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		fmt.Println("Error parsing proxy")
		return err
	}
	if err != nil {
		fmt.Println("error setting proxy")
		return err
	}
	Transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: Transport}
	pl.client = client
	//	pl.currentProxy = New(proxyStr)
	pl.currentProxy = pl.Pool[ind]
	pl.currentInd = ind
	return nil
}

// Returns currently set proxy
func (pl *ProxyList) CurrentProxy() (*Proxy, error) {
	pr := pl.currentProxy
	fmt.Println(pr)
	if pr.p == "" {
		err := errCurrentProxyUnset
		return nil, err
	}
	return pr, nil
}

func (pl *ProxyList) DeleteProxy() error {
	if pl.currentProxy.p == "" {
		err := errCurrentProxyUnset
		return err
	}
	ind := pl.currentInd
	pl.Pool = append(pl.Pool[:ind], pl.Pool[ind+1:]...)
	return nil
}

// Get method with timeout
func (pl *ProxyList) Get(url string, t time.Duration) (*http.Response, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, t*time.Second)
	defer cancel()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	res, err := pl.client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
