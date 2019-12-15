package recursion

func validTicTacToe(board []string) bool {
	dimension := len(board)
	//to record X and O in rows and cols
	rows, cols := make([][2]int, dimension), make([][2]int, dimension)
	//to record X and O in diagnals
	xdiag1, xdiag2, odiag1, odiag2 := 0, 0, 0, 0
	//to record number of appearance of X and O
	xnum, onum := 0, 0

	for i, r := range board {
		for j, c := range r {
			if c == 'X' {
				xnum += 1
				rows[i][0] += 1
				cols[j][0] += 1
				if i == j {
					xdiag1 += 1
				}
				if i+j == dimension-1 {
					xdiag2 += 1
				}
			} else if c == 'O' {
				onum += 1
				rows[i][1] += 1
				cols[j][1] += 1
				if i == j {
					odiag1 += 1
				}
				if i+j == dimension-1 {
					odiag2 += 1
				}
			}
		}
	}

	xwin := false
	owin := false
	//check X win or O win
	for _, r := range rows {
		if r[0] == 3 {
			xwin = true
		} else if r[1] == 3 {
			owin = true
		}
	}
	//check X win or O win
	for _, c := range cols {
		if c[0] == 3 {
			xwin = true
		} else if c[1] == 3 {
			owin = true
		}
	}
	//check X win or O win
	if xdiag1 == 3 || xdiag2 == 3 {
		xwin = true
	}
	if odiag1 == 3 || odiag2 == 3 {
		owin = true
	}

	//x and o could not be simutaneously win
	if xwin && owin {
		return false
	}
	//if X win, then difference of X and O should be 1
	if xwin {
		if xnum-onum == 1 {
			return true
		} else {
			return false
		}
	}

	//if  O win, then difference of X and O should be 0
	if owin {
		if onum == xnum {
			return true
		} else {
			return false
		}
	}


	//if none of them win, then, the difference should be 0 or 1
	if xnum-onum == 0 || xnum-onum == 1 {
		return true
	} else {
		return false
	}
}
