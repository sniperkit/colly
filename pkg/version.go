package colly

// Must be set at build via
// -ldflags "-X main.VERSION=`cat VERSION`"
// -ldflags "-X main.VERSION=`git describe --tags`"
var VERSION string
