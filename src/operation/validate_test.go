package operation

import (
	"testing"
)

func TestSquareMatrix(t *testing.T) {
	t.Run("return error if number of cols are greater than rows", func(t *testing.T) {
		matrix := [][]string{
			{"1","2","3"},
			{"4","5","6"},
		}
		err := squareMatrix(matrix)
		assertError(t, err, errInvalidMatrix)
	})
	t.Run("return error if number of rows are greater than cols", func(t *testing.T) {
		matrix := [][]string{
			{"1","2"},
			{"3","4"},
			{"5","6"},
		}
		err := squareMatrix(matrix)
		assertError(t, err, errInvalidMatrix)
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{},
		}
		err := squareMatrix(matrix)
		assertError(t, err, errInvalidMatrix)
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{},
		}
		err := squareMatrix(matrix)
		if err != errInvalidMatrix {
			t.Fatal("didn't get invalid matrix error but wanted it")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{"j"},
		}
		_, err := convertMatrixToInt(matrix)
		if err != errInvalidMatrix {
			t.Fatal("didn't get invalid matrix error but wanted it")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{""},
		}
		_, err := convertMatrixToInt(matrix)
		if err != errInvalidMatrix {
			t.Fatal("didn't get invalid matrix error but wanted it")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{" "},
		}
		_, err := convertMatrixToInt(matrix)
		if err != errInvalidMatrix {
			t.Fatal("didn't get invalid matrix error but wanted it")
		}
	})
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
