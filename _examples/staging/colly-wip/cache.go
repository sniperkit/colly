package main

import (
	"encoding/csv"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	// cache - core
	"github.com/gregjones/httpcache"

	// cache - backends
	"github.com/gregjones/httpcache/memcache"
	/*
		"github.com/birkelund/boltdbcache"
		"github.com/klaidliadon/go-couch-cache"
		"github.com/klaidliadon/go-memcached"
		"github.com/klaidliadon/go-redis-cache"
		"github.com/gregjones/httpcache/diskcache"
		"sourcegraph.com/sourcegraph/s3cache"
	*/

	// cache - advanced backends
	"github.com/sniperkit/xcache/backend/default/badger"
	"github.com/sniperkit/xcache/backend/default/diskv"

	// general helpers
	"github.com/sniperkit/xtask/util/fs"
)

/*
	Refs:
	- https://github.com/docker/leeroy/blob/master/github/github.go
	- https://github.com/calavera/openlandings/blob/master/github/transport.go
	- https://github.com/Dreae/esi-graphql/blob/master/resolvers/http/init.go
*/

// defaults
var (
	defaultCacheEngine       string        = "inmemory"
	defaultCacheFreqUnit     string        = "Week"
	defaultCacheFreqDuration time.Duration = 1
	defaultCacheTTL          time.Duration = time.Duration(defaultCacheFreqDuration * time.Hour)
	defaultCacheBackends     []string      = []string{"inmemory", "redis", "sqlite3", "badger", "mysql", "postgres"}
)

// cache related objects
var (
	xClient       *http.Client
	xTransport    *httpcache.Transport
	xCache        httpcache.Cache
	xRoundTripper http.RoundTripper
)

// global cache variables
var (
	isCacheTransport  bool = true
	cacheVary         bool = true
	cacheForce        bool = false
	cacheStatus       bool = false
	cacheError        error
	cacheTimeDuration int           = 1
	cacheTimeUnit     string        = "Week"
	cacheTTL          time.Duration = time.Duration(defaultCacheFreqDuration * time.Hour)
	cacheErrors       []error       = make([]error, 0)
)

// cache backend variables
var (
	cacheEngine      = "badger"
	cachePrefixPath  = ""
	cacheStoragePath = cachePrefixPath + "./shared/storage/cache/http"
)

func init() {
	sm.CreateDirs(defaultStorageDirs)
}

func cloneCacheHTTP() httpcache.Cache {
	defer funcTrack(time.Now())
	backendCache, err := newCacheBackend(cacheEngine, cacheStoragePath)
	if err != nil {
		log.Fatal("cache err", err.Error())
	}
	return backendCache
}

func newCacheWithTransport(engine string, prefixPath string) (httpcache.Cache, *httpcache.Transport) {
	cacheBackend, err := newCacheBackend(engine, prefixPath)
	if err != nil {
		log.Fatalln("error:", err)
	}

	cacheTransport := httpcache.NewTransport(cacheBackend)
	cacheTransport.MarkCachedResponses = true
	// cacheRoundTripper := newCacheRoundTripper(cacheTransport, true)
	return cacheBackend, cacheTransport // , cacheRoundTripper
}

func newCacheBackend(engine string, prefixPath string) (backend httpcache.Cache, err error) {
	defer funcTrack(time.Now())
	fsutil.EnsureDir(prefixPath)
	engine = strings.ToLower(engine)
	switch cacheEngine {
	case "diskv":
		cacheStoragePath = filepath.Join(prefixPath, "cacher.diskv")
		fsutil.EnsureDir(cacheStoragePath)
		backend = diskcache.New(cacheStoragePath)
	case "badger":
		cacheStoragePath = filepath.Join(prefixPath, "cacher.badger")
		fsutil.EnsureDir(cacheStoragePath)
		backend, err = badgercache.New(
			&badgercache.Config{
				ValueDir:    "golanglibs.com",
				StoragePath: cacheStoragePath,
				SyncWrites:  false,
				Debug:       false,
				Compress:    true,
				TTL:         time.Duration(120 * 24 * time.Hour),
			},
		)
	case "memory":
		backend = httpcache.NewMemoryCache()
	default:
		err = errInvalidCacheBackend
	}
	return
}

func newCacheRoundTripper(c httpcache.Cache, markCachedResponses bool) http.RoundTripper {
	defer funcTrack(time.Now())

	t := httpcache.NewTransport(c)
	t.MarkCachedResponses = markCachedResponses
	return t
}

// InitHTTP initializes the HTTP client using an appropriate cache service
// ref.
func initClientHTTP() {
	if memcachedURL := os.Getenv("MEMCACHE_URL"); memcachedURL != "" {
		xClient = httpcache.NewTransport(memcache.New(memcachedURL)).Client()
	} else {
		xClient = httpcache.NewTransport(httpcache.NewMemoryCache()).Client()
	}
}

func setGroupCache(createdAt time.Time, group string, key string, obj map[string]interface{}) {
	xCache.Set(key, toBytes(mapToString(obj)))
}

func setCache(key string, obj map[string]interface{}) {
	xCache.Set(key, toBytes(mapToString(obj)))
}

func isValidBackendTTL(b string, d time.Duration, u string) (ttl time.Duration, ok bool, err error) {
	if d <= 0 {
		err = errInvalidCacheDuration
		return
	}
	u = strings.ToLower(u)
	switch u {
	case "week":
		ttl = time.Duration(d * time.Hour * 24 * 7)
	case "day":
		ttl = time.Duration(d * time.Hour * 24)
	case "hour":
		ttl = time.Duration(d * time.Hour)
	case "second":
		ttl = time.Duration(d * time.Second)
	default:
		err = errInvalidCacheTimeUnit
		return
	}
	if ttl <= -1 {
		err = errInvalidCacheTTL
		return
	}
	// start cache storage backend
	ok = true
	return
}

func loadSkipListFromCache(filepath string) {
	fp, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	csv := csv.NewReader(fp)
	lines := streamCsv(csv, csvReaderBuffer)
	for line := range lines {
		expiresAt, err := parseTimeStamp(line.GetByName("task_expired_timestamp"))
		if err != nil {
			log.Errorln("[SKIP-ERROR] taskInfo, service=", line.GetByName("service"), "topic=", line.GetByName("topic"), "expiresTimestamp", line.GetByName("task_expired_timestamp"))
			continue
		}
		now := time.Now()
		if now.After(expiresAt.Add(cacheTTL)) {
			log.Infoln("[TSK-ALLOW] task info, service=", line.GetByName("service"), "topic=", line.GetByName("topic"), "expiresAt=", expiresAt)
			continue
		}
		cuckflt.InsertUnique([]byte(line.GetByName("topic")))
	}
	log.Warnln("[TSK-EXCLUDED] taskInfo, count=", cuckflt.Count())
}