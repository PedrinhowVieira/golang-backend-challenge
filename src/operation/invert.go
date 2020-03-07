package operation

import (
	"strconv"
)

// Invert takes a matrix of string elements, invert the rows an columns
// and returns it as a single string formatted as a matrix.
func Invert(matrix [][]string) string {
	if squareMatrix(matrix) != nil {
		return stringNonSquareMatrix + "\n"
	}
	return invertConcurrency(matrix)
}

func invert(matrix [][]string) string {
	invertedMatrix := copyMatrix(matrix)
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			_, err := strconv.Atoi(matrix[i][j])
			if err != nil {
				return notInteger
			}
			invertedMatrix[j][i] = matrix[i][j]
		}
	}
	response := Echo(invertedMatrix)
	return response
}

func invertConcurrency(matrix [][]string) string {
	_, intErr := convertMatrixToInt(matrix)
	if intErr != nil {
		return notInteger
	}

	invertedMatrix := copyMatrix(matrix)
	channel := make(chan [][]string)

	for i, _ := range matrix {
		go func(v int) {
			for j, _ := range matrix[v] {
				invertedMatrix[j][v] = matrix[v][j]
			}
			channel <- invertedMatrix
		}(i)
	}
	for i := 0; i < len(invertedMatrix); i++ {
		invertedMatrix = <-channel
	}
	response := echo(invertedMatrix)
	return response
}