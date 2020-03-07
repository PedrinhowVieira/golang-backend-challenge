package operation

import (
	"fmt"
	"strings"
)

// Flatten takes a matrix of string elements and returns it as a single string.
func Flatten(matrix [][]string) string {
	if squareMatrix(matrix) != nil {
		return stringNonSquareMatrix
	}
	_, intErr := convertMatrixToInt(matrix)
	if intErr != nil {
		return notInteger
	}

	return flatten(matrix)
}

func flatten(matrix [][]string) string {
	responses := make([]string, len(matrix))
	for i, row := range matrix {
		responses[i] = fmt.Sprintf("%s%s", responses[i], strings.Join(row, ","))
	}
	response := strings.Join(responses[:], ",") + "\n"
	return response
}