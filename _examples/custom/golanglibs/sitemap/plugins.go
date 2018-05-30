package main

import (
	"plugin"
)

var (
	defaultCollectorPluginFilepath string = "./plugins/bitcq/lib/bitcq.so"
)

func loadPlugin(filePath string) {
	p, err := plugin.Open(filePath)
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("Search")
	if err != nil {
		panic(err)
	}

	search := f.(func(string, string) []byte)
	results := search("Unfriended", "qwerty")
	fmt.Println(string(results))
}