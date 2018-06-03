package config

import (
	"regexp"
)

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
///// Collector's filters registry
//////////////////////////////////////////////////

type Filters struct {
	Disabled        bool            `default:'false' json:'enabled' yaml:'enabled' toml:'enabled' xml:'enabled' ini:'enabled'`
	WhiteListRules  []*FilterConfig `json:'whitelist_rules' yaml:'whitelist_rules' toml:'whitelist_rules' xml:'whitelistRules' ini:'whitelistRules'`
	BlackListRules  []*FilterConfig `json:'blacklist_rules' yaml:'blacklist_rules' toml:'blacklist_rules' xml:'blackListRules' ini:'blackListRules'`
	WatchListRules  []*FilterConfig `json:'watchlist_rules' yaml:'watchlist_rules' toml:'watchlist_rules' xml:'watchListRules' ini:'watchListRules'`
	NotifyListRules []*FilterConfig `json:'notifylist_rules' yaml:'notifylist_rules' toml:'notifylist_rules' xml:'notifyListRules' ini:'notifyListRules'`
	errs            []*error        `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
}

func (fs *Filters) IsEnabled() bool {
	return fs.enabled
}

func (fs *Filters) Enable() {
	fs.Enabled = true
}

func (fs *Filters) Disable() {
	fs.Enabled = false
}

func (fs *Filters) ListRules(filterBy string) (errs []string) {
	for _, e := range fs.rules {
		errs = append(errs, e.Error())
	}
	return
}

func (fs *Filters) Status() bool {
	return fs.Enabled
}

func (fs *Filters) String() string {
	return ""
}

func (fs *Filters) Errors(level string) []*error {
	return fs.errs
}

func (fs *Filters) ListErrors(level string) (errs []string) {
	for _, e := range fs.errs {
		errs = append(errs, e.Error())
	}
	return
}

//////////////////////////////////////////////////
///// Collector's filter actions
//////////////////////////////////////////////////

type FilterConfig struct {
	Disabled    bool       `default:'false' json:'enabled' yaml:'enabled' toml:'enabled' xml:'enabled' ini:'enabled'`
	Rule        string     `default:'' json:'rule' yaml:'rule' toml:'rule' xml:'rule' ini:'rule'`
	ScannerType filterType `default:'regexp' json:'scanner_type' yaml:'scanner_type' toml:'scanner_type' xml:'ScannerType' ini:'ScannerType'`
	isValid     bool
}

func (f *Filter) String() string {
	return ""
}

func (f *Filter) IsValid(pattern string) bool {
	return f.isValid
}

func (f *Filter) AddRule(pattern string) bool {
	return false
}

func (f *Filter) IsValid(pattern string) bool {
	return false
}

func checkRuleByName(pattern string) bool {
	return false
}