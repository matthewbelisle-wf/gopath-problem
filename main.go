package main

import (
	"fmt"
	"bar"
	"net/http"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bar: %d", bar.Bar)
	})
}
