package main

import (
	"fmt"
	"net/http"

	"echo"
	"readfile"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		records := readfile.ReadFile(w, r)
		result := echo.Echo(records)
		fmt.Fprint(w, result)
	})
	http.HandleFunc("/invert", func(w http.ResponseWriter, r *http.Request) {
		records := readfile.ReadFile(w, r)
		result := echo.Echo(records)
		fmt.Fprint(w, result)
	})
	http.ListenAndServe(":8080", nil)
}
