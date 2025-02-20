package operation

import (
	"strconv"
)

var invertBenchmarkThreshold = 350

// Invert takes a matrix of string elements, invert the rows an columns
// and returns it as a single string formatted as a matrix.
func Invert(matrix [][]string) string {
	// analyses the matrix size to choose the algorithm with most performance based on a benchmark test
	if  len(matrix) < invertBenchmarkThreshold {
		return invertDefault(matrix)
	} else {
		return invertConcurrency(matrix)
	}
}

func invertDefault(matrix [][]string) string {
	invertedMatrix := copyMatrix(matrix)
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			invertedMatrix[j][i] = matrix[i][j]
		}
	}
	response := Echo(invertedMatrix)
	return response
}

func invertConcurrency(matrix [][]string) string {
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
	response := Echo(invertedMatrix)
	return response
}

func invertConversionIn(matrix [][]string) string {
	invertedMatrix := copyMatrix(matrix)
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			_, err := strconv.Atoi(matrix[i][j])
			if err != nil {
				return InvalidMatrix
			}
			invertedMatrix[j][i] = matrix[i][j]
		}
	}
	response := Echo(invertedMatrix)
	return response
}
