package operation

import (
	"errors"
)

var notInteger = "cannot insert a non-integer value\n"
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