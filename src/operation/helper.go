package operation

import (
	"errors"
	"strconv"
)

var notInteger = "cannot insert a non-integer value"
var errNotInteger = errors.New(notInteger)

func convertToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return i, errNotInteger
	}
	return i, nil
}

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