package validate

import (
	"errors"
)

var StringNonSquareMatrix = "cannot insert a non-square matrix"
var ErrNonSquareMatrix = errors.New(StringNonSquareMatrix)

func SquareMatrix(matrix [][]string) error {
	rows := len(matrix)
	cols := len(matrix[rows - 1])
	if rows == 1 || rows != cols{
		return ErrNonSquareMatrix
	}
	return nil
}