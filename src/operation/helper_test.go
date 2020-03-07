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
		assertError(t, err, errNonSquareMatrix)
	})
	t.Run("return error if number of rows are greater than cols", func(t *testing.T) {
		matrix := [][]string{
			{"1","2"},
			{"3","4"},
			{"5","6"},
		}
		err := squareMatrix(matrix)
		assertError(t, err, errNonSquareMatrix)
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{},
		}
		err := squareMatrix(matrix)
		assertError(t, err, errNonSquareMatrix)
	})
	t.Run("do nothing if the number of rows and cols are the same", func(t *testing.T) {
		matrix := [][]string{
			{"1","2"},
			{"3","4"},
		}
		err := squareMatrix(matrix)
		if err != nil {
			t.Errorf("should not returned an error but it did")
		}
	})
}

func TestConvertToInt(t *testing.T) {
	t.Run("takes a string and returns an integer", func(t *testing.T) {
		got, _ := convertToInt("8")
		want := 8
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("returns an error when the input is not a integer", func(t *testing.T) {
		_, err := convertToInt("j")
		assertNotIntError(t, err, errNotInteger)
	})
	t.Run("returns an error when the input is empty", func(t *testing.T) {
		_, err := convertToInt("")
		assertNotIntError(t, err, errNotInteger)
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

func assertNotIntError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}