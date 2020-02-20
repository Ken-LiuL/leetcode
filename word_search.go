package array

var R int
var C int
func exist(board [][]byte, word string) bool {
	R, C = len(board), len(board[0])
    for i := 0; i < R; i++ {
        for j := 0; j < C; j++ {
            if helper(i, j , board, word) {
                return true
            } 
        }
    }
    return false
}

func helper(r, c int, board [][]byte, word string) bool {
	  if word == "" {
		  return true
	  }
	  if r < 0 || r > R-1 || c < 0 || c > C -1 {
		  return false
	  }
	  if board[r][c] == 0 || board[r][c] != word[0] {
		  return false
	  }
	  rest , v := word[1:], board[r][c]
	  board[r][c] = 0
	  res := helper(r+1, c, board, rest) || helper(r-1, c, board, rest) || 
			 helper(r, c+1, board, rest) || helper(r, c-1, board, rest)
      if !res {
		  board[r][c] = v
	  }
	  return res
	  
}