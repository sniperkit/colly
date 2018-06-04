package config

type ClientConfig struct {
	// Enabled
	Enabled bool `default:"false" json:"enabled" yaml:"enabled" toml:"enabled" xml:"enabled" ini:"enabled" csv:"Enabled"`

	// Address
	Address string `default:"" json:"address" yaml:"address" toml:"address" xml:"address" ini:"address" csv:"address"`

	// Domain
	Domain string `json:"domain" yaml:"domain" toml:"domain" xml:"domain" ini:"domain"`

	// Protocol
	Protocol string `json:"protocol" yaml:"protocol" toml:"protocol" xml:"protocol" ini:"protocol"`

	// Host
	Host string `json:"host" yaml:"host" toml:"host" xml:"host" ini:"host"`

	// Port
	Port string `json:"port" yaml:"port" toml:"port" xml:"port" ini:"port"`

	// ForceSSL
	ForceSSL bool `default:"true" json:"force_ssl" yaml:"force_ssl" toml:"force_ssl" xml:"force_ssl" ini:"force_ssl"`

	// VerifySSL
	VerifySSL bool `default:"false" json:"ssl_verify" yaml:"ssl_verify" toml:"ssl_verify" xml:"verifySSL" ini:"verifySSL"`

	// Reconnect
	Reconnect bool `default:"true" json:"reconnect" yaml:"reconnect" toml:"reconnect" xml:"reconnect" ini:"reconnect"`

	// ReconnectMax
	ReconnectMax int `default:"3" json:"reconnect_max" yaml:"reconnect_max" toml:"reconnect_max" xml:"reconnectMax" ini:"reconnectMax"`

	// BufferSize
	BufferSize string `json:"buffer_size" yaml:"buffer_size" toml:"buffer_size" xml:"bufferSize" ini:"bufferSize"`

	// Payload
	Payload string `json:"payload" yaml:"payload" toml:"payload" xml:"payload" ini:"payload"`

	//-- End
}

type ProxyConfig struct {

	// Enabled
	Enabled bool `default:"true" json:"enabled" yaml:"enabled" toml:"enabled" xml:"enabled" ini:"enabled" csv:"Enabled"`

	// Address
	Address string `required:"true" json:"address" yaml:"address" toml:"address" xml:"address" ini:"address" csv:"address"`

	//-- End
}
