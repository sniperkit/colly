package config

import (
	"time"
)

const (
	DEFAULT_STORAGE_CACHE_DIR string = "./shared/cache/queue/internal"
)

// defaults
var (
	defaultCacheEngine       string        = "inmemory"
	defaultCacheFreqUnit     string        = "Week"
	defaultCacheFreqDuration time.Duration = 1
	defaultCacheTTL          time.Duration = time.Duration(defaultCacheFreqDuration * time.Hour)
	defaultCacheBackends     []string      = []string{"inmemory", "redis", "sqlite3", "badger", "mysql", "postgres"}
)

var (
	defaultDomain                string   = "https://golanglibs\\.com"
	defaultStorageExportDir      string   = "./shared/storage/export"
	defaultStorageCacheDir       string   = "./shared/storage/cache"
	defaultStorageLogDir         string   = "./shared/logs"
	defaultStorageSitemapDirname string   = defaultStorageCacheDir + "/sitemaps"
	defaultStorageDirs           []string = []string{
		defaultStorageCacheDir,
		defaultStorageLogDir,
		defaultStorageExportDir,
		defaultStorageSitemapDirname,
	}
)
