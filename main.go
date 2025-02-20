package main

import (
	"file"
	"fmt"
	"net/http"
	"operation"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func Handler(w http.ResponseWriter, r *http.Request) {
	matrix := file.Read(w, r)
	intMatrix, err := operation.Validate(matrix)
	if err != nil {
		fmt.Fprint(w, operation.InvalidMatrix)
	} else {
		var result string
		switch r.URL.Path {
		case "/echo":
			result = operation.Echo(matrix)
		case "/flatten":
			result = operation.Flatten(matrix)
		case "/invert":
			result = operation.Invert(matrix)
		case "/multiply":
			result = operation.Multiply(intMatrix)
			result += "\n"
		case "/sum":
			result = operation.Sum(intMatrix)
			result += "\n"
		default:
			result = "Error: wrong path\n"
		}
		fmt.Fprint(w, result)
	}
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
