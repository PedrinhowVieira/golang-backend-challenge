package operation

import (
"testing"
)

func TestFlatten(t *testing.T) {
	t.Run("flat a 2x2 matrix to string", func(t *testing.T) {
		matrix := [][]string{
			{"2","2"},
			{"2","2"},
		}
		got := Flatten(matrix)
		want := "2,2,2,2\n"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("flat any matrix size to string", func(t *testing.T) {
		matrix := [][]string{
			{"1","2","3"},
			{"4","5","6"},
			{"7","8","9"},
		}
		got := Flatten(matrix)
		want := "1,2,3,4,5,6,7,8,9\n"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
