package operation

import (
	"math/rand"
	"strconv"
)

func copyMatrix(matrix [][]string) [][]string{
	size := len(matrix)
	copiedMatrix := make([][]string, size)
	for i, _ := range matrix {
		copiedMatrix[i] = make([]string, size)
	}
	return copiedMatrix
}

func generateMatrix(size int) [][]string {
	matrix := make([][]string, size)
	for i, _ := range matrix {
		matrix[i] = make([]string, size)
	}
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			matrix[i][j] = strconv.Itoa(rand.Intn(100))
		}
	}
	return matrix
}
