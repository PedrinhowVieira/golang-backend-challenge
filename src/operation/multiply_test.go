package operation

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	t.Run("multiply all elements of the a 2x2 matrix with values of 2", func(t *testing.T) {
		matrix := [][]int{
			{2,2},
			{2,2},
		}
		got := Multiply(matrix)
		want := "16"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("multiply all elements of any matrix size with any values", func(t *testing.T) {
		matrix := [][]int{
			{1,2,3},
			{4,5,6},
			{7,8,9},
		}
		got := Multiply(matrix)
		want := "362880"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("multiply all elements of the a 2x2 matrix with values of 2, with concurrency", func(t *testing.T) {
		matrix := [][]int{
			{2,2},
			{2,2},
		}
		got := multiplyConcurrency(matrix)
		want := "16"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("multiply all elements of any matrix size with any values, with concurrency", func(t *testing.T) {
		matrix := [][]int{
			{1,2,3},
			{4,5,6},
			{7,8,9},
		}
		got := multiplyConcurrency(matrix)
		want := "362880"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}

func BenchmarkMultiplyDefault(b *testing.B) {
	intMatrix, _ := generateMatrix(sumBenchmarkThreshold)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = multiplyDefault(intMatrix)
	}
}
func BenchmarkMultiplyConcurrency(b *testing.B) {
	intMatrix, _ := generateMatrix(sumBenchmarkThreshold)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = multiplyConcurrency(intMatrix)
	}
}
func BenchmarkMultiplyConversionIn(b *testing.B) {
	_, matrix := generateMatrix(multiplyBenchmarkThreshold)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = multiplyConversionIn(matrix)
	}
}
