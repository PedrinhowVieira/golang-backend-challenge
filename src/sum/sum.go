package sum

import (
	"strconv"
)

// Sum takes a matrix of string elements and return the sum of all elements as a string
func Sum(matrix [][]string) int {
	var sum int
		for i, _ := range matrix {
			for j, _ := range matrix[i] {
				sum += convertToInt(matrix[i][j])
			}
		}
	return sum
}

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
