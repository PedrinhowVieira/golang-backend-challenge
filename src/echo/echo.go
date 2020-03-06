package echo

import (
	"fmt"
	"strings"
	"validate"
)

// Echo returns the matrix as a string in matrix format.
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
