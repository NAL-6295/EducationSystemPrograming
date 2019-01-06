package main

import (
	"encoding/csv"
	"net/http"
	"os"
)

func main() {

	request, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "へっだーもついかできます")
	csvWriter := csv.NewWriter(os.Stdout)
	csvWriter.Write([]string{"test", "test"})
	csvWriter.Write([]string{"test", "test2"})
	csvWriter.Write([]string{request.URL.Path})
	csvWriter.Flush()

}
