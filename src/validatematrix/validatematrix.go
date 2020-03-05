package validatematrix

import "errors"

var ErrNonSquareMatrix = errors.New("cannot insert a non-square matrix")

func ValidateSquareMatrix(matrix [][]string) error {

	return ErrNonSquareMatrix
}