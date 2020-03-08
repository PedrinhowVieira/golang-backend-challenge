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
		want := "8"
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
		want := "45"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("return error if number of cols are greater than rows", func(t *testing.T) {
		matrix := [][]string{
			{"1","2","3"},
			{"4","5","6"},
		}
		err := Sum(matrix)
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
		err := Sum(matrix)
		if err != stringNonSquareMatrix {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{},
		}
		err := Sum(matrix)
		if err != stringNonSquareMatrix {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{"j"},
		}
		err := Sum(matrix)
		if err != notInteger {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{""},
		}
		err := Sum(matrix)
		if err != notInteger {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{" "},
		}
		err := Sum(matrix)
		if err != notInteger {
			t.Fatal("didn't get an error but wanted one")
		}
	})
}

func BenchmarkSumDefault(b *testing.B) {
	matrix := generateMatrix(sumBenchmarkThreshold)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sumDefault(matrix)
	}
}
func BenchmarkSumConcurrency(b *testing.B) {
	matrix := generateMatrix(sumBenchmarkThreshold)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sumConcurrency(matrix)
	}
}
func BenchmarkSumConversionIn(b *testing.B) {
	matrix := generateMatrix(sumBenchmarkThreshold)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sumConversionIn(matrix)
	}
}
