package operation

import "strconv"

// Multiply takes a matrix of string elements and return the multiplication of all elements as a string.
func Multiply(matrix [][]string) string {
	if squareMatrix(matrix) != nil {
		return stringNonSquareMatrix + "\n"
	}

	total := 1
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			value, err := convertToInt(matrix[i][j])
			if err != nil {
				return notInteger + "\n"
			}
			total *= value
		}
	}
	response := strconv.Itoa(total) + "\n"
	return response
}
