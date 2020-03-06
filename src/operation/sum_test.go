package operation

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum all elements of the a 2x2 matrix with values of 2", func(t *testing.T) {
		matrix := [][]string{
			{"2","2"},
			{"2","2"},
		}
		got := Sum(matrix)
		want := "8\n"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("sum all elements of any matrix size with any values", func(t *testing.T) {
		matrix := [][]string{
			{"1","2","3"},
			{"4","5","6"},
			{"7","8","9"},
		}
		got := Sum(matrix)
		want := "45\n"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
