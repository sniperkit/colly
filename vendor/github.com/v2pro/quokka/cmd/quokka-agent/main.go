package main

import (
	"net/http"
	"github.com/v2pro/quokka/agent"
)

func main() {
	http.ListenAndServe(":8318", agent.Mux)
}