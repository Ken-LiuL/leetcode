package dynamic

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func findLength(A []int, B []int) int {
	m, n := len(A)+1, len(B)+1
	maxLength := 0
    now := make([]int, n)
	for i := 1; i < m; i++ {
		next := make([]int, n)
		for j := 1; j < n; j++ {
			if A[i-1] == B[j-1] {
				next[j] = now[j-1] + 1
				maxLength = max(next[j], maxLength)
			}
		}
		now = next
	}
	return maxLength

}