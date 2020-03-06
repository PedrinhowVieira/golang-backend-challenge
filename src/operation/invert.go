package operation

// Echo takes a matrix of string elements and returns it as a single string formatted as a matrix.
func Invert(matrix [][]string) string {
	size := len(matrix)
	invertedMatrix := make([][]string, size)
	for i, _ := range matrix {
		invertedMatrix[i] = make([]string, size)
	}
	for i, row := range matrix {
		for j, col := range row {
			invertedMatrix[j][i] = col
		}
	}
	response := Echo(invertedMatrix)
	return response
}
