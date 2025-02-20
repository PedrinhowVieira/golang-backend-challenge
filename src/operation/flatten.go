package operation

import (
	"fmt"
	"strings"
)

// Flatten takes a matrix of string elements and returns it as a single string.
func Flatten(matrix [][]string) string {
	responses := make([]string, len(matrix))
	for i, row := range matrix {
		responses[i] = fmt.Sprintf("%s%s", responses[i], strings.Join(row, ","))
	}
	response := strings.Join(responses[:], ",") + "\n"
	return response
}
