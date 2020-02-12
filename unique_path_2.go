package dynamic 

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	for r := m -1 ; r >= 0; r-- {
		for c := n -1; c >= 0; c-- {
			if obstacleGrid[r][c] == 1 {
				obstacleGrid[r][c] = 0
			} else {
				if 	r == m-1 && c == n - 1 {
					obstacleGrid[r][c] = 1
				} else if r == m-1 {
					obstacleGrid[r][c] = obstacleGrid[r][c+1]
				} else if c == n -1 {
					obstacleGrid[r][c] = obstacleGrid[r+1][c]
				} else {
					obstacleGrid[r][c] = obstacleGrid[r][c+1] + obstacleGrid[r+1][c]
				}
			}
		}
	}
	return obstacleGrid[0][0]
}

	