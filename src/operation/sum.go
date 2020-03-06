package operation

import "strconv"

// Sum takes a matrix of string elements and return the sum of all elements as a string
func Sum(matrix [][]string) string {
	var sum int
		for i, _ := range matrix {
			for j, _ := range matrix[i] {
				sum += convertToInt(matrix[i][j])
			}
		}
	response := strconv.Itoa(sum) + "\n"
	return response
}

