# colly-sitemap-create

## Purpose

Crawl a website a generate a sitemap in xml or txt format. (Optional: can generate a sitemap index if found more than 10.000 entries)


## Requirements

go 1.8+

## Installation

```bash
go get -t -v github.com/sniperkit/colly/examples/experimental/sitemap-create
```

## Compilation

Inside your GOPATH directory

```bash
cd src/github.com/sniperkit/colly/examples/experimental/sitemap-create
go build -buildmode=plugin -o plugins/sitemap-create/sitemap-create.so main.go
go build -o colly-sitemap-create main.go
```

## Testing

```bash
go test -v ./...
```
