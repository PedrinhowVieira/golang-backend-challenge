package operation

import (
	"testing"
)

func TestEcho(t *testing.T) {
	t.Run("convert a 2x2 matrix to string", func(t *testing.T) {
		matrix := [][]string{
			{"2","2"},
			{"2","2"},
		}
		got := Echo(matrix)
		want := "2,2\n2,2\n"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("convert any matrix size to string", func(t *testing.T) {
		matrix := [][]string{
			{"1","2","3"},
			{"4","5","6"},
			{"7","8","9"},
		}
		got := Echo(matrix)
		want := "1,2,3\n4,5,6\n7,8,9\n"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("return error if number of cols are greater than rows", func(t *testing.T) {
		matrix := [][]string{
			{"1","2","3"},
			{"4","5","6"},
		}
		err := Echo(matrix)
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
		err := Echo(matrix)
		if err != stringNonSquareMatrix+"\n" {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{},
		}
		err := Echo(matrix)
		if err != stringNonSquareMatrix+"\n" {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{"j"},
		}
		err := Echo(matrix)
		if err != notInteger {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{""},
		}
		err := Echo(matrix)
		if err != notInteger {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{" "},
		}
		err := Echo(matrix)
		if err != notInteger {
			t.Fatal("didn't get an error but wanted one")
		}
	})
}
