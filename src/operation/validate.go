package operation

import (
	"errors"
	"strconv"
)

var InvalidMatrix = "the inserted matrix is invalid\n"
var errInvalidMatrix = errors.New(InvalidMatrix)

func Validate(matrix [][]string) ([][]int, error) {
	var intMatrix [][]int
	var intErr error
	intMatrix, intErr = convertMatrixToInt(matrix)
	if squareMatrix(matrix) != nil || intErr != nil {
		return intMatrix, errInvalidMatrix
	} else {
		return intMatrix, nil
	}
}

func squareMatrix(matrix [][]string) error {
	rows := len(matrix)
	if rows == 0 {
		return errInvalidMatrix
	}
	cols := len(matrix[rows - 1])
	if rows != cols {
		return errInvalidMatrix
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
				return convertedMatrix, errInvalidMatrix
			}
			convertedMatrix[i][j] = value
		}
	}
	return convertedMatrix, nil
}