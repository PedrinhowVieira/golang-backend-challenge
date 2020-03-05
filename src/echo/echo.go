package echo

import (
	"fmt"
	"strings"
)

// Echo returns the matrix as a string in matrix format.
func Echo(records [][]string) string {
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}
