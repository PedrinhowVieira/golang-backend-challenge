package operation

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum all elements of the a 2x2 matrix with values of 2", func(t *testing.T) {
		matrix := [][]int{
			{2,2},
			{2,2},
		}
		got := Sum(matrix)
		want := "8"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("sum all elements of any matrix size with any values", func(t *testing.T) {
		matrix := [][]int{
			{1,2,3},
			{4,5,6},
			{7,8,9},
		}
		got := Sum(matrix)
		want := "45"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("sum all elements of the a 2x2 matrix with values of 2, with concurrency", func(t *testing.T) {
		matrix := [][]int{
			{2,2},
			{2,2},
		}
		got := sumConcurrency(matrix)
		want := "8"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("sum all elements of any matrix size with any values, with concurrency", func(t *testing.T) {
		matrix := [][]int{
			{1,2,3},
			{4,5,6},
			{7,8,9},
		}
		got := sumConcurrency(matrix)
		want := "45"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}

func BenchmarkSumDefault(b *testing.B) {
	intMatrix, _ := generateMatrix(sumBenchmarkThreshold)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sumDefault(intMatrix)
	}
}
func BenchmarkSumConcurrency(b *testing.B) {
	intMatrix, _ := generateMatrix(sumBenchmarkThreshold)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sumConcurrency(intMatrix)
	}
}
func BenchmarkSumConversionIn(b *testing.B) {
	_, matrix := generateMatrix(sumBenchmarkThreshold)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sumConversionIn(matrix)
	}
}
