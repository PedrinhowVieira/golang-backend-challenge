package operation

import "strconv"

// Multiply takes a matrix of string elements and return the multiplication of all elements as a string.
func Multiply(matrix [][]string) string {
	value := 1
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			value *= convertToInt(matrix[i][j])
		}
	}
	response := strconv.Itoa(value) + "\n"
	return response
}
