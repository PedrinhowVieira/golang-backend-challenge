package main

import (
	"fmt"
	"net/http"
	"operation"
	"readfile"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func Handler(w http.ResponseWriter, r *http.Request) {
	records := readfile.ReadFile(w, r)
	result := ""
	switch r.URL.Path {
	case "/echo":
		result = operation.Echo(records)
	case "/sum":
		result = operation.Sum(records)
	case "/multiply":
		result = operation.Multiply(records)
	default:
		result = "Error: wrong path\n"
	}
	fmt.Fprint(w, result)
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
