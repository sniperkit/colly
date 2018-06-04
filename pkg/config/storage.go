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

type StoreConfig struct {
	Enabled   bool   `default:"false" json:"enabled" yaml:"enabled" toml:"enabled" xml:"enabled" ini:"enabled" csv:"enabled"`
	Domain    string `json:'domain' yaml:'domain' toml:'domain' xml:'domain' ini:'domain'`
	Protocol  string `json:'protocol' yaml:'protocol' toml:'protocol' xml:'protocol' ini:'protocol'`
	Host      string `json:'host' yaml:'host' toml:'host' xml:'host' ini:'host'`
	Port      string `json:'port' yaml:'port' toml:'port' xml:'port' ini:'port'`
	ForceSSL  bool   `default:'true' json:'force_ssl' yaml:'force_ssl' toml:'force_ssl' xml:'force_ssl' ini:'force_ssl'`
	VerifySSL bool   `default:'false' json:'ssl_verify' yaml:'ssl_verify' toml:'ssl_verify' xml:'verifySSL' ini:'verifySSL'`
}
