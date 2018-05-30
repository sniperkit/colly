package colly

// RequestCallback is a type alias for OnRequest callback functions
type RequestCallback func(*Request)

// ResponseCallback is a type alias for OnResponse callback functions
type ResponseCallback func(*Response)

// HTMLCallback is a type alias for OnHTML callback functions
type HTMLCallback func(*HTMLElement)

// TABCallback is a type alias for OnTAB callback functions
type RAWCallback func(*RAWElement)

// BINCallback is a type alias for OnTAB callback functions
// type BINCallback func(*BINElement)

// XMLCallback is a type alias for OnXML callback functions
type XMLCallback func(*XMLElement)

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

type rawCallbackContainer struct {
	Selector string
	Function RAWCallback
}
