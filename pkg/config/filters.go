package config

var (
	// Default allowed domains
	DefaultAllowedDomains []string = []string{
		"golanglibs.com", "github.com", "gitlab.com", "bitbucket.com", "stackoverflow.com", "reddit.com", "medium.com",
	}
	// Default disabled domains
	DefaultDisallowedDomains []string = []string{
		"google.com", "bing.com", "yahoo.com",
	}
)

// Backend represents a KV Store Backend
type filterType string

const (
	REGEXP filterType = "regexp" // Using default golang regexp package (default)
	PLUCK  filterType = "pluck"  // Alternative to xpath and regexp, it allows to extract pattern with activators/desactivators patterns
	LEXER  filterType = "lexer"  // to do...
	RUNE   filterType = "rune"   // to do...
	XQUERY filterType = "xquery" // to do...
	XPATH  filterType = "xpath"  // to do...
	AST    filterType = "ast"    // to do...
)

//////////////////////////////////////////////////
///// Collector"s filters registry
//////////////////////////////////////////////////

type Filters struct {
	Disabled        bool           `default:"false" json:"enabled" yaml:"enabled" toml:"enabled" xml:"enabled" ini:"enabled"`
	WhiteListRules  []FilterConfig `json:"whitelist_rules" yaml:"whitelist_rules" toml:"whitelist_rules" xml:"whitelistRules" ini:"whitelistRules"`
	BlackListRules  []FilterConfig `json:"blacklist_rules" yaml:"blacklist_rules" toml:"blacklist_rules" xml:"blackListRules" ini:"blackListRules"`
	WatchListRules  []FilterConfig `json:"watchlist_rules" yaml:"watchlist_rules" toml:"watchlist_rules" xml:"watchListRules" ini:"watchListRules"`
	NotifyListRules []FilterConfig `json:"notifylist_rules" yaml:"notifylist_rules" toml:"notifylist_rules" xml:"notifyListRules" ini:"notifyListRules"`
	errs            []*error       `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-"`
}

func (fs *Filters) IsEnabled() bool {
	return fs.Disabled
}

func (fs *Filters) Enable() {
	fs.Disabled = false
}

func (fs *Filters) Disable() {
	fs.Disabled = true
}

func (fs *Filters) ListRules(filterBy string) (rules []string) {
	//for _, e := range fs.WhiteListRules {
	//	errs = append(errs, e.Error())
	//}
	return
}

func (fs *Filters) Status() bool {
	return fs.Disabled
}

func (fs *Filters) String() string {
	return ""
}

func (fs *Filters) Errors(level string) []*error {
	return fs.errs
}

func (fs *Filters) ListErrors(level string) (errs []string) {
	//for _, e := range fs.errs {
	//	errs = append(errs, e)
	//}
	return
}

//////////////////////////////////////////////////
///// Collector"s filter actions
//////////////////////////////////////////////////

type DomainFilter struct {
	Disabled  bool   `default:"false" json:"disabled" yaml:"disabled" toml:"disabled" xml:"disabled" ini:"disabled"`
	Domain    string `json:"domain" yaml:"domain" toml:"domain" xml:"domain" ini:"domain"`
	Protocol  string `json:"protocol" yaml:"protocol" toml:"protocol" xml:"protocol" ini:"protocol"`
	Host      string `json:"host" yaml:"host" toml:"host" xml:"host" ini:"host"`
	Port      string `json:"port" yaml:"port" toml:"port" xml:"port" ini:"port"`
	ForceSSL  bool   `default:"true" json:"force_ssl" yaml:"force_ssl" toml:"force_ssl" xml:"force_ssl" ini:"force_ssl"`
	VerifySSL bool   `default:"false" json:"ssl_verify" yaml:"ssl_verify" toml:"ssl_verify" xml:"verifySSL" ini:"verifySSL"`
	up        bool   `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-"`
	healthy   bool   `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-"`
	err       error  `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-"`
}

//////////////////////////////////////////////////
///// Collector"s filter actions
//////////////////////////////////////////////////

type FilterConfig struct {
	Enabled     bool   `default:"true" json:"enabled" yaml:"enabled" toml:"enabled" xml:"enabled" ini:"enabled"`
	Rule        string `default:"" json:"rule" yaml:"rule" toml:"rule" xml:"rule" ini:"rule"`
	ScannerType string `default:"regex" json:"scanner" yaml:"scanner" toml:"scanner" xml:"scanner" ini:"scanner"`
	isValid     bool   `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-"`
}

func (f *FilterConfig) String() string {
	return ""
}

func (f *FilterConfig) IsValid(pattern string) bool {
	return f.isValid
}

func (f *FilterConfig) AddRule(pattern string) bool {
	return false
}

func checkRuleByName(pattern string) bool {
	return false
}
