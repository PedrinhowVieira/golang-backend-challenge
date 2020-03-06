package echo

import (
	"fmt"
	"strings"
	"validate"
)

// Echo takes a matrix of string elements and returns it as a single string formatted as a matrix.
func Echo(matrix [][]string) string {
	if validate.SquareMatrix(matrix) != nil {
		return validate.StringNonSquareMatrix + "\n"
	}
	var response string
	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}
