package operation

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	t.Run("multiply all elements of the a 2x2 matrix with values of 2", func(t *testing.T) {
		matrix := [][]string{
			{"2","2"},
			{"2","2"},
		}
		got := Multiply(matrix)
		want := "16\n"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("multiply all elements of any matrix size with any values", func(t *testing.T) {
		matrix := [][]string{
			{"1","2","3"},
			{"4","5","6"},
			{"7","8","9"},
		}
		got := Multiply(matrix)
		want := "362880\n"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("return error if number of cols are greater than rows", func(t *testing.T) {
		matrix := [][]string{
			{"1","2","3"},
			{"4","5","6"},
		}
		err := Multiply(matrix)
		if err != stringNonSquareMatrix+"\n" {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error if number of rows are greater than cols", func(t *testing.T) {
		matrix := [][]string{
			{"1","2"},
			{"3","4"},
			{"5","6"},
		}
		err := Multiply(matrix)
		if err != stringNonSquareMatrix+"\n" {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{},
		}
		err := Multiply(matrix)
		if err != stringNonSquareMatrix+"\n" {
			t.Fatal("didn't get an error but wanted one")
		}
	})
}
