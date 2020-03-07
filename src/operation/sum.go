package operation

import "strconv"

// Sum takes a matrix of string elements and return the sum of all elements as a string.
func Sum(matrix [][]string) string {
	if squareMatrix(matrix) != nil {
		return stringNonSquareMatrix + "\n"
	}
	return sum(matrix)
}

func sum(matrix [][]string) string {
	var total int
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			value, err := strconv.Atoi(matrix[i][j])
			if err != nil {
				return notInteger
			}
			total += value
		}
	}
	response := strconv.Itoa(total) + "\n"
	return response
}

