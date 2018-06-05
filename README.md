# What
A fully resumable horizontally (infinitely) scalable webscraper in Go. Think 1000's of machines scraping sites in a distributed way. Based on Docker swarm, Cassandra, Traefik, colly, gRPC, and my other [boilerplate](https://github.com/dioptre/gtrpc).

# Note
You could probably just use colly... Especially if you don't care about scalability... or use a shell script ([Example](https://github.com/dioptre/scrp/blob/master/simple.sh))

# Why
I built this to distribute scraping across multiple servers, so as to go undetected. I could have used proxies, but wanted to reuse the code for other distributed apps.

# Instructions
*Read* and run build.sh for instructions on building and executing. Just edit /service/scrape.go to customize what you want to upload to Cassandra and how. Then run the client (code in /client/client.go for details). For example:
```
./gcli https://en.wikipedia.org/wiki/List_of_HTTP_status_codes _ ".*wikipedia\.org.*"
```

# Thanks
Cheers to the engineers of Cassandra, colly, gRPC,Consul, Traefik & protobuf to name a few.

