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
	matrix := readfile.ReadFile(w, r)
	result := ""
	switch r.URL.Path {
	case "/echo":
		result = operation.Echo(matrix)
	case "/sum":
		result = operation.Sum(matrix)
	case "/multiply":
		result = operation.Multiply(matrix)
	case "/flatten":
		result = operation.Flatten(matrix)
	default:
		result = "Error: wrong path\n"
	}
	fmt.Fprint(w, result)
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
