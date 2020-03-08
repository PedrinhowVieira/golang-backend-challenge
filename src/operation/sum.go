package operation

import (
	"strconv"
)

var sumBenchmarkThreshold = 2500

// Sum takes a matrix of string elements and return the sum of all elements as a string.
func Sum(matrix [][]string) string {
	if squareMatrix(matrix) != nil {
		return stringNonSquareMatrix
	}
	// analyses the matrix size to choose the algorithm with most performance based on a benchmark test
	if len(matrix) < sumBenchmarkThreshold {
		return sumDefault(matrix)
	} else {
		return sumConcurrency(matrix)
	}
}

func sumDefault(matrix [][]string) string {
	intMatrix, intErr := convertMatrixToInt(matrix)
	if intErr != nil {
		return notInteger
	}

	var total int
	for i, _ := range intMatrix {
		for j, _ := range intMatrix[i] {
			total += intMatrix[i][j]
		}
	}
	response := strconv.Itoa(total)
	return response
}

func sumConcurrency(matrix [][]string) string {
	intMatrix, intErr := convertMatrixToInt(matrix)
	if intErr != nil {
		return notInteger
	}
	channel := make(chan int)

	for i, _ := range intMatrix {
		go func(v int) {
			var total int
			for j, _ := range intMatrix[v] {
				total += intMatrix[v][j]
			}
			channel <- total
		}(i)
	}
	var total int
	for i := 0; i < len(intMatrix); i++ {
		total += <-channel
	}
	response := strconv.Itoa(total)
	return response
}

func sumConversionIn(matrix [][]string) string {
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
	response := strconv.Itoa(total)
	return response
}
