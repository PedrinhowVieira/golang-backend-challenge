package operation

import (
	"errors"
	"math/rand"
	"strconv"
)

var notInteger = "cannot insert a non-integer value\n"
var errNotInteger = errors.New(notInteger)
var stringNonSquareMatrix = "cannot insert a non-square matrix\n"
var errNonSquareMatrix = errors.New(stringNonSquareMatrix)
var benchmarkMatrixSize = 2000

func squareMatrix(matrix [][]string) error {
	rows := len(matrix)
	if rows == 0 {
		return errNonSquareMatrix
	}
	cols := len(matrix[rows - 1])
	if rows != cols {
		return errNonSquareMatrix
	}
	return nil
}

func convertMatrixToInt(matrix [][]string) ([][]int, error) {
	size := len(matrix)
	convertedMatrix := make([][]int, size)
	for i, _ := range matrix {
		convertedMatrix[i] = make([]int, size)
	}
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			value, err := strconv.Atoi(matrix[i][j])
			if err != nil {
				return convertedMatrix, errNotInteger
			}
			convertedMatrix[i][j] = value
		}
	}
	return convertedMatrix, nil
}

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
