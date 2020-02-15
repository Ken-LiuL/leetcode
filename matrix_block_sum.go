package dynamic

func matrixBlockSum(mat [][]int, K int) [][]int {
	var r, c int = len(mat), len(mat[0])
	var res [][]int = make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	for i := 0; i <= K; i++ {
		for j := 0; j <= K; j++ {
			if i < r && j < c {
				res[0][0] += mat[i][j]
			}

		}
	}

	for j := 1; j < c; j++ {
		var tmp int = res[0][j-1]
		for t := 0; t <= K; t++ {
			if j-1-K >= 0 && t < r {
				tmp -= mat[t][j-1-K]
			}
			if j+K < c && t < r {
				tmp += mat[t][j+K]
			}
		}
		res[0][j] = tmp
	}

	for i := 1; i < r; i++ {
		res[i][0] = res[i-1][0]
		for t := 0; t <= K; t++ {
			if t < c && i-1-K >= 0 {
				res[i][0] -= mat[i-1-K][t]
			}
			if t < c && i+K < r {
				res[i][0] += mat[i+K][t]
			}
		}
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			res[i][j] = res[i][j-1] + res[i-1][j] - res[i-1][j-1]
			if i+K < r && j+K < c {
				res[i][j] += mat[i+K][j+K]
			}
			if i-1-K >= 0 && j-1-K >= 0 {
				res[i][j] += mat[i-1-K][j-1-K]
			}
			if i+K < r && j-1-K >= 0 {
				res[i][j] -= mat[i+K][j-1-K]
			}
			if j+K < c && i-1-K >= 0 {
				res[i][j] -= mat[i-1-K][j+K]
			}
		}
	}
	return res
}