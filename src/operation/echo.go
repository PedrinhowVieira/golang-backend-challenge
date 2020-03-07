package operation

import (
	"fmt"
	"strings"
)

// Echo takes a matrix of string elements and returns it as a single string formatted as a matrix.
func Echo(matrix [][]string) string {
	if squareMatrix(matrix) != nil {
		return stringNonSquareMatrix + "\n"
	}
	_, intErr := convertMatrixToInt(matrix)
	if intErr != nil {
		return notInteger
	}

	return echo(matrix)
}

func echo(matrix [][]string) string {
	var response string
	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}
