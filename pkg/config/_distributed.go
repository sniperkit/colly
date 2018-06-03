package config

/*
import (
	"time"
)

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

type Sep struct {
	Rune rune
}

func (r *Sep) UnmarshalText(text []byte) error {
	if len(text) > 0 {
		data := int32(text[0])
		r.Rune = data
	}
	return nil
}

type DistributedConfig struct {
	expiredAt             duration           `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
	blockSize             int32              `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
	servers               []string           `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
	server                string             `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
	port                  string             `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
	pemFile               string             `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
	statusCodesAcceptance map[string]float64 `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
	retryOnStatusCodes    []int              `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
	retryCount            int                `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
	ignoreAttrs           []string           `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
	isTLSMode             bool               `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
}

//////////////////////////////////////////////////
///// Getters
//////////////////////////////////////////////////
func (cfgDist *Distributed) GetTimeout(format *string) time.Duration {
	return cfgDist.expiredAt
}

func (cfgDist *Distributed) GetBlockSize() int32 {
	return cfgDist.blockSize
}

//////////////////////////////////////////////////
///// Setters
//////////////////////////////////////////////////
func (cfgDist *Distributed) SetTimeout(expireAt time.Duration) {
	cfgDist.expiredAt = expireAt
	return
}

func (cfgDist *Distributed) WithTimeout(expireAt time.Duration) *Distributed {
	cfgDist.expiredAt = expireAt
	return cfgDist.expiredAt
}

func (cfgDist *Distributed) SetBlockSize(blockSize int32) {
	cfgDist.blockSize = blockSize
	return
}

func (cfgDist *Distributed) WithBlockSize(blockSize int32) *Distributed {
	cfgDist.blockSize = blockSize
	return cfgDist.blockSize
}

// NB. create others methods...
*/
