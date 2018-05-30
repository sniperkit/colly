package main

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gregjones/httpcache"
	cuckoo "github.com/seiflotfy/cuckoofilter"
	"github.com/willf/bloom"

	"github.com/sniperkit/xcache/backend/default/badger"
	"github.com/sniperkit/xcache/backend/default/diskv"
	"github.com/sniperkit/xtask/util/fs" // move into a separate repo/package
)

var xcache httpcache.Cache

// defaults
var (
	defaultCacheEngine       string        = "inmemory"
	defaultCacheFreqUnit     string        = "Week"
	defaultCacheFreqDuration time.Duration = 1
	defaultCacheTTL          time.Duration = time.Duration(defaultCacheFreqDuration * time.Hour)
	defaultCacheBackends     []string      = []string{"inmemory", "redis", "sqlite3", "badger", "mysql", "postgres"}
)

// global cache variables
var (
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

// cache lists - bloom filters
var (
	bloomFilterSize uint = 20000 // default: 500000
	bloomFilterKeys uint = 5
	blmflt               = bloom.New(bloomFilterSize, bloomFilterKeys)
)

// cache lists - cuckoo filters
var (
	cuckFilterCapacity uint = 20000 // default: 1000000
	cuckFilter              = cuckoo.NewCuckooFilter(cuckFilterCapacity)
	cuckflt                 = cuckoo.NewDefaultCuckooFilter()
)

func initCacheHTTP(b string, d time.Duration, u string) (ttl time.Duration, ok bool, err error) {
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

func cloneCacheHTTP() httpcache.Cache {
	defer funcTrack(time.Now())
	backendCache, err := newCacheHTTP(cacheEngine, cacheStoragePath)
	if err != nil {
		log.Fatal("cache err", err.Error())
	}
	return backendCache
}

func newCacheHTTP(engine string, prefixPath string) (backend httpcache.Cache, err error) {
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
				ValueDir:    "api.github.com.v3.snappy",
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

func setServiceCache(createdAt time.Time, service string, key string, obj map[string]interface{}) {
	xcache.Set(key, toBytes(mapToString(obj)))
}

func setCache(key string, obj map[string]interface{}) {
	xcache.Set(key, toBytes(mapToString(obj)))
}

func loadSkipList(filepath string) {
	fp, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	csv := csv.NewReader(fp)
	lines := streamCsv(csv, csvReaderBuffer)
	for line := range lines {
		expiresAt, err := parseTimeStamp(line.Get("task_expired_timestamp"))
		if err != nil {
			log.Errorln("[SKIP-ERROR] taskInfo, service=", line.Get("service"), "topic=", line.Get("topic"), "expiresTimestamp", line.Get("task_expired_timestamp"))
			continue
		}
		now := time.Now()
		if now.After(expiresAt.Add(cacheTTL)) {
			log.Infoln("[TSK-ALLOW] task info, service=", line.Get("service"), "topic=", line.Get("topic"), "expiresAt=", expiresAt)
			continue
		}
		cuckflt.InsertUnique([]byte(line.Get("topic")))
	}
	log.Warnln("[TSK-EXCLUDED] taskInfo, count=", cuckflt.Count())
}
