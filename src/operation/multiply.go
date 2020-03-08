package operation

import "strconv"

var multiplyBenchmarkThreshold = 2500

// Multiply takes a matrix of string elements and return the multiplication of all elements as a string.
func Multiply(matrix [][]string) string {
	if squareMatrix(matrix) != nil {
		return stringNonSquareMatrix
	}
	// analyses the matrix size to choose the algorithm with most performance based on a benchmark test
	if len(matrix) < multiplyBenchmarkThreshold {
		return multiplyDefault(matrix)
	} else {
		return multiplyConcurrency(matrix)
	}
}

func multiplyDefault(matrix [][]string) string {
	intMatrix, intErr := convertMatrixToInt(matrix)
	if intErr != nil {
		return notInteger
	}
	total := 1
	for i, _ := range intMatrix {
		for j, _ := range intMatrix[i] {
			total *= intMatrix[i][j]
		}
	}
	response := strconv.Itoa(total)
	return response
}

func multiplyConcurrency(matrix [][]string) string {
	intMatrix, intErr := convertMatrixToInt(matrix)
	if intErr != nil {
		return notInteger
	}
	channel := make(chan int)

	for i, _ := range intMatrix {
		go func(v int) {
			total := 1
			for j, _ := range intMatrix[v] {
				total *= intMatrix[v][j]
			}
			channel <- total
		}(i)
	}
	total := 1
	for i := 0; i < len(intMatrix); i++ {
		total *= <-channel
	}
	response := strconv.Itoa(total)
	return response
}

func multiplyConversionIn(matrix [][]string) string {
	total := 1
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			value, err := strconv.Atoi(matrix[i][j])
			if err != nil {
				return notInteger
			}
			total *= value
		}
	}
	response := strconv.Itoa(total)
	return response
}
