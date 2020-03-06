package validate

import (
	"testing"
)

func TestSquareMatrix(t *testing.T) {
	t.Run("return error if number of cols are greater than rows", func(t *testing.T) {
		matrix := [][]string{
			{"1","2","3"},
			{"4","5","6"},
		}
		err := SquareMatrix(matrix)
		assertError(t, err, ErrNonSquareMatrix)
	})
	t.Run("return error if number of rows are greater than cols", func(t *testing.T) {
		matrix := [][]string{
			{"1","2"},
			{"3","4"},
			{"5","6"},
		}
		err := SquareMatrix(matrix)
		assertError(t, err, ErrNonSquareMatrix)
	})
	t.Run("return error if it is a single number matrix", func(t *testing.T) {
		matrix := [][]string{
			{"1"},
		}
		err := SquareMatrix(matrix)
		assertError(t, err, ErrNonSquareMatrix)
	})
	t.Run("do nothing if the number of rows and cols are the same", func(t *testing.T) {
		matrix := [][]string{
			{"1","2"},
			{"3","4"},
		}
		err := SquareMatrix(matrix)
		if err != nil {
			t.Errorf("should not returned an error but it did")
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