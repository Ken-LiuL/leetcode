package array


func transpose(A [][]int) [][]int {
	c, r := len(A), len(A[0])
	newM := make([][]int, r)
	for i := 0; i < r; i++ {
		newM[i] = make([]int, c)
		for j := 0; j < c; j++ {
			newM[i][j] = A[j][i]
		}
	}
	return newM
}