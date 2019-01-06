package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World"
	}
}