package operation

import (
	"errors"
	"strconv"
)

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

var stringNonSquareMatrix = "cannot insert a non-square matrix"
var errNonSquareMatrix = errors.New(stringNonSquareMatrix)

func squareMatrix(matrix [][]string) error {
	rows := len(matrix)
	cols := len(matrix[rows - 1])
	if rows != cols{
		return errNonSquareMatrix
	}
	return nil
}