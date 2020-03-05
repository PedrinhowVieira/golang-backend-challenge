package validatematrix

import (
"testing"
)

func TestValidateSquareMatrix(t *testing.T) {
	t.Run("validate if the number of rows and columns are the same", func(t *testing.T) {
		matrix := [][]string{
			{"1"},
			{"3","4"},
		}
		err := ValidateSquareMatrix(matrix)
		assertError(t, err, ErrNonSquareMatrix)
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