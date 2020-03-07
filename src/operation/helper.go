package operation

import (
	"errors"
	"strconv"
)

var notInteger = "cannot insert a non-integer value\n"
var errNotInteger = errors.New(notInteger)
var stringNonSquareMatrix = "cannot insert a non-square matrix"
var errNonSquareMatrix = errors.New(stringNonSquareMatrix)

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