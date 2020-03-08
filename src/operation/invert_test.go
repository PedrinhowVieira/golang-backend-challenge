package operation

import (
	"testing"
)

func TestInvert(t *testing.T) {
	t.Run("invert the rows and columns a 2x2 matrix to string", func(t *testing.T) {
		matrix := [][]string{
			{"1","2"},
			{"3","4"},
		}
		got := Invert(matrix)
		want := "1,3\n2,4\n"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("invert the rows and columns of any matrix size to string", func(t *testing.T) {
		matrix := [][]string{
			{"1","2","3"},
			{"4","5","6"},
			{"7","8","9"},
		}
		got := Invert(matrix)
		want := "1,4,7\n2,5,8\n3,6,9\n"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("return error if number of cols are greater than rows", func(t *testing.T) {
		matrix := [][]string{
			{"1","2","3"},
			{"4","5","6"},
		}
		err := Invert(matrix)
		if err != stringNonSquareMatrix {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error if number of rows are greater than cols", func(t *testing.T) {
		matrix := [][]string{
			{"1","2"},
			{"3","4"},
			{"5","6"},
		}
		err := Invert(matrix)
		if err != stringNonSquareMatrix {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{},
		}
		err := Invert(matrix)
		if err != stringNonSquareMatrix {
			t.Fatal("didn't get an error but wanted one")
		}
	})

	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{"j"},
		}
		err := Invert(matrix)
		if err != notInteger {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{""},
		}
		err := Invert(matrix)
		if err != notInteger {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{" "},
		}
		err := Invert(matrix)
		if err != notInteger {
			t.Fatal("didn't get an error but wanted one")
		}
	})
}

func BenchmarkInvertDefault(b *testing.B) {
	matrix := generateMatrix(invertBenchmarkThreshold)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = invertDefault(matrix)
	}
}
func BenchmarkInvertConcurrency(b *testing.B) {
	matrix := generateMatrix(invertBenchmarkThreshold)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = invertConcurrency(matrix)
	}
}
func BenchmarkInvertConversionIn(b *testing.B) {
	matrix := generateMatrix(invertBenchmarkThreshold)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = invertConversionIn(matrix)
	}
}
