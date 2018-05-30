# check-my-pages

**check-my-pages** is a scrapping script to test large sites. We are using it to test the Spanish soon-to-be archived site and redirects system, with more than 10,000 pages. It checks redirections, http responses, analytics, files hosted in soon-to-die servers, canonical urls and more.

 It checks each url in a list and creates report files about what was tested. Each file reports about a specific issue and includes the scanned url together with the result.

Please note that the script was written to check for problems detected when manually inspecting a small sample (about 1%) of the Spanish website. ~~This script was not built to be used as-it-is in other websites, you are supposed to modify it.~~

This **command line** script complements other command line tools like ack, grep, rpl and others. Our own [ecounter](https://github.com/greenpeace/gpes-ecounter) and [ecompare](https://github.com/greenpeace/gpes-ecompare) are also used.

## Download and install

### Install the lastest version

Go to the [releases page](https://github.com/greenpeace/gpes-check-my-pages/releases) and download the last version for your 64bit operating system: Windows, Mac and Linux.

### Install from source

If you have the [Go](https://golang.org/) compiler installed, you can download and install go with:

```bash
go get github.com/greenpeace/gpes-check-my-pages

go install github.com/greenpeace/gpes-check-my-pages
```

### Update from source

```bash
go get -u github.com/greenpeace/gpes-check-my-pages

go install github.com/greenpeace/gpes-check-my-pages
```

### Get help

If you downloaded and installed the latest version from the [releases page](https://github.com/greenpeace/gpes-check-my-pages/releases) do:

```bash
 ./check-my-pages --help
```

If you installed from source do:

```bash
 check-my-pages --help
```

## File with list of urls

The urls file, by default `urls.csv` must have all the urls you want to check. You can use a text file with 1 url per line or a csv file with the urls on the first column and without headers.

You can use [ecounter](https://github.com/greenpeace/gpes-ecounter) to create a urls file from a sitemap.xml file.

## Http info about a list of urls

If you want to obtain information about http status codes, mime-types, file sizes and redirect urls of any urls, you can use `-http`.

You must use this check in a separate command like:

```bash
./check-my-pages -urls=urls.csv -http -miliseconds=100
```

because check-my-pages will stop after executing `-http`

This check creates a file named `httpResponses.csv` with 5 fields: 

1. initial url
1. http status code
1. mime type
1. file size *(adds -1 if the file size is unknown)*
1. final url

## Checking html urls

To do all the checks in `urls.csv` (html urls) with all the checks use the command:

```bash
./check-my-pages -urls=urls.csv -analytics -canonical -title -linkpattern -cssjspattern -mediapattern
```

This repository includes a few testing urls in the file `urls.csv`. Please replace them by your own.

It will create a couple of files, one per check the script is doing:

* `analytics.csv` - Reports **google analytics** tracking ID
* `canonicals.csv` - Reports the **canonical url** for every url
* `titles.csv` - Reports the **title** for every url
* `linkpattern.csv` - Reports on links that include a regular expression pattern. Useful to track **links** to specific **dead sites**. The default pattern can be set by the `-pattern` option.
* `cssjspattern.csv` - Reports **css and js** urls that include a regular expression pattern. To detect dead css and js urls in large sites. The pattern can also be defined with the option `-pattern` (described bellow)
* `mediapattern.csv` - Reports **media** links. Images, videos, audios, iframes and objects. Also use `-pattern` to define the urls pattern.

## Optional command line configurations

`-miliseconds=100` - Sets a delay of 100 miliseconds between requests (the default value)

`-pattern='https?://(\w|-)+.greenpeace.org/espana/.+'` - Changes the search url pattern to the regular expression.

## Remove the report files

To remove the files created by **check-my-pages**:

```bash
./check-my-pages -clear
```

## Rename the report files

Renames the files created by **check-my-pages** adding date and time to the name.

```bash
./check-my-pages -stash
```

By renaming the files you prevent the results from future reports being added to the end of each file.

## Crawl URLS

If you don't have a sitemap.xml or another file with the urls, you can crawl the site: 

```bash
./check-my-pages -crawl -urls=crawledurls.csv  -start='https://www.fotografar.net/' -pattern='https://www.fotografar.net/.*' -miliseconds=100
```

It will save, in crawledurls.csv, all the urls that match the regular expression pattern and that can be obtained by crawling from the start url. 

You can make as many crawls as you want to the same crawledurls.csv file. If you get repeated urls, you can use [ecounter](https://github.com/greenpeace/gpes-ecounter) to obtain unique urls (from one or multiple concatenated files).
