package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "world",
	}

	g := gzip.NewWriter(w)
	mw := io.MultiWriter(g, os.Stdout)

	e := json.NewEncoder(mw)

	e.Encode(source)
	g.Flush()

	a := make([]int, 5)
	b := append(a[:1], 5, 4)

	fmt.Println(b)

}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
