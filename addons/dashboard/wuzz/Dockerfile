FROM golang:1.9-alpine

ENV GOPATH $HOME/gocode
ENV PATH $GOPATH/bin:$PATH

RUN apk add --no-cache \
      git

COPY docker/entry.sh /docker-entrypoint

RUN chmod u+x /docker-entrypoint && \
	go get github.com/asciimoo/wuzz

ENTRYPOINT ["/docker-entrypoint"]