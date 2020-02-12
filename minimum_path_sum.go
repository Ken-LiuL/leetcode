package dynamic


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for r := m - 1; r >=0; r-- {
		for c := n - 1; c >=0; c-- {
			if r == m - 1 && c == n - 1 {
				grid[r][c] = grid[r][c]
			} else if r == m - 1 {
				grid[r][c] += grid[r][c+1]
			} else if c == n - 1 {
				grid[r][c] += grid[r+1][c]
			} else {
				grid[r][c] += min(grid[r][c+1], grid[r+1][c])
			}
		}
	}
	return grid[0][0]
}