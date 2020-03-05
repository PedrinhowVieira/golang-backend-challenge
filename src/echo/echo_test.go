package echo

import (
	"testing"
)

func TestEcho(t *testing.T) {
	t.Run("convert matrix to string", func(t *testing.T) {
		matrix := [][]string{
			{"1","2"},
			{"3","4"},
		}
		got := Echo(matrix)
		want := "1,2\n3,4\n"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
