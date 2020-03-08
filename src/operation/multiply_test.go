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
		want := "16"
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
		want := "362880"
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
		err := Multiply(matrix)
		if err != stringNonSquareMatrix {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{},
		}
		err := Multiply(matrix)
		if err != stringNonSquareMatrix {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{"j"},
		}
		err := Multiply(matrix)
		if err != notInteger {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{""},
		}
		err := Multiply(matrix)
		if err != notInteger {
			t.Fatal("didn't get an error but wanted one")
		}
	})
	t.Run("return error there is no number inside the matrix", func(t *testing.T) {
		matrix := [][]string{
			{" "},
		}
		err := Multiply(matrix)
		if err != notInteger {
			t.Fatal("didn't get an error but wanted one")
		}
	})
}

//func BenchmarkMultiplyDefault(b *testing.B) {
//	matrix := generateMatrix(benchmarkMatrixSize)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		_ = multiplyDefault(matrix)
//	}
//}
//func BenchmarkMultiplyDefault2(b *testing.B) {
//	matrix := generateMatrix(benchmarkMatrixSize)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		_ = multiplyDefault2(matrix)
//	}
//}
//func BenchmarkMultiplyConcurrency(b *testing.B) {
//	matrix := generateMatrix(benchmarkMatrixSize)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		_ = multiplyConcurrency(matrix)
//	}
//}
