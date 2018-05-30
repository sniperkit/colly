# Makefile - X-Colly

EXAMPLE_GOLANGLIBS_SRC := "./_examples/custom/golanglibs/sitemap"
EXAMPLE_GOLANGLIBS_BIN := "$(CURDIR)/bin/golanglibs"

golanglib-build:
	@rm -fR $(EXAMPLE_GOLANGLIBS_BIN)
	@go build -o $(EXAMPLE_GOLANGLIBS_BIN) $(EXAMPLE_GOLANGLIBS_SRC)
	@$(EXAMPLE_GOLANGLIBS_BIN) --help


build-plugins:
	@go build -buildmode=plugin -o ./shared/libs/bitcq.so ./addons/plugins/bitcq/main.go