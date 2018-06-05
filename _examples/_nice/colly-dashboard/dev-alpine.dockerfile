FROM golang:1.10.2-alpine3.7 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/sniperkit/colly/_examples/_nice/colly_dashboard

# COPY vendor ../vendor
# COPY bin ./
RUN go build -o /go/bin/app main.go

FROM alpine:3.7
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]