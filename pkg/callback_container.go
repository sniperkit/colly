package colly

// htmlCallbackContainer
type htmlCallbackContainer struct {
	Selector string       `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
	Function HTMLCallback `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
}

// xmlCallbackContainer
type xmlCallbackContainer struct {
	Query    string      `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
	Function XMLCallback `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
}

// jsonCallbackContainer
type jsonCallbackContainer struct {
	Query    string       `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
	Function JSONCallback `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
}

// tabCallbackContainer
type tabCallbackContainer struct {
	Query    string      `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
	Hook     *TABHook    `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
	Function TABCallback `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
}
