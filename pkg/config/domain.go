package config

type DomainConfig struct {
	Disabled  bool   `default:'false' json:'disabled' yaml:'disabled' toml:'disabled' xml:'disabled' ini:'disabled'`
	Domain    string `json:'domain' yaml:'domain' toml:'domain' xml:'domain' ini:'domain'`
	Protocol  string `json:'protocol' yaml:'protocol' toml:'protocol' xml:'protocol' ini:'protocol'`
	Host      string `json:'host' yaml:'host' toml:'host' xml:'host' ini:'host'`
	Port      string `json:'port' yaml:'port' toml:'port' xml:'port' ini:'port'`
	ForceSSL  bool   `default:'true' json:'force_ssl' yaml:'force_ssl' toml:'force_ssl' xml:'force_ssl' ini:'force_ssl'`
	VerifySSL bool   `default:'false' json:'ssl_verify' yaml:'ssl_verify' toml:'ssl_verify' xml:'verifySSL' ini:'verifySSL'`
	up        bool   `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
	healthy   bool   `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
	err       error  `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
}
