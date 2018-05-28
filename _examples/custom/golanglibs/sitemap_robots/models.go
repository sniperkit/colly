package main

type repo struct {
	Author  string `selector:"a.hnuser"`
	URL     string `selector:".age a[href]" attr:"href"`
	Comment string `selector:".comment"`
	depth   int
}

type similar struct {
	Author  string `selector:"a.hnuser"`
	URL     string `selector:".age a[href]" attr:"href"`
	Comment string `selector:".comment"`
	depth   int
}

type event struct {
	Author  string `selector:"a.hnuser"`
	URL     string `selector:".age a[href]" attr:"href"`
	Comment string `selector:".comment"`
	depth   int
}
