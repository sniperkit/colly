package colly

// RequestCallback is a type alias for OnRequest callback functions
type RequestCallback func(*Request)

// ResponseCallback is a type alias for OnResponse callback functions
type ResponseCallback func(*Response)

// HTMLCallback is a type alias for OnHTML callback functions
type HTMLCallback func(*HTMLElement)

// XMLCallback is a type alias for OnXML callback functions
type XMLCallback func(*XMLElement)

// JSONCallback is a type alias for OnJSON callback functions
type JSONCallback func(*JSONElement)

// CollectorCallback is a type alias for OnEvent callback functions
type CollectorCallback func(*Collector)

// FuncCallback is a type alias for OnFunc callback functions
type FuncCallback func(*Collector)

// EventCallback is a type alias for OnEvent callback functions
type EventCallback func(*Collector)

// ErrorCallback is a type alias for OnError callback functions
type ErrorCallback func(*Response, error)

// ScrapedCallback is a type alias for OnScraped callback functions
type ScrapedCallback func(*Response)

type htmlCallbackContainer struct {
	Selector string
	Function HTMLCallback
}

type xmlCallbackContainer struct {
	Query    string
	Function XMLCallback
}

type jsonCallbackContainer struct {
	Query    string
	Function JSONCallback
}
