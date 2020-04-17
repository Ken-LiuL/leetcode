package UnionFind

type UnionFind struct {
	arr  []int
	sz   []int
}

func NewUnionFind(size int) *UnionFind {
   uf := &UnionFind{arr: make([]int, size), sz: make([]int, size)}
   return uf
}

func (uf *UnionFind) root(p int) int {
	for p != uf.arr[p] {
		p = uf.arr[p]
	}
	return p
}
func (uf *UnionFind) union(p, q int) {
   rp, rq := uf.root(p), uf.root(q)
	if rp == rq {
		 return 
	} else { 
		 sp, sq := uf.sz[rp], uf.sz[rq]
		 if sp < sq {
			uf.arr[rp] = rq
			uf.sz[rq] += uf.sz[rp]
		 } else {
			uf.arr[rq]  = rp
			uf.sz[rp] += uf.sz[rq] 
		 }
	}
}
func (uf *UnionFind) connect(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

func solve(board [][]byte)  {
	if len(board) == 0 {
        return
    }
	//init union find and with additional element as super parent 
	//of all nodes in the edge
	row_num, col_num := len(board), len(board[0])
	//craete additional element for super parent
	uf := NewUnionFind(row_num*col_num+1)
	super_parent :=  row_num*col_num
	//init union find array
	for  rindex, r := range board {
		for cindex, c := range r {
			if c == 'O' {
				id := rindex*col_num+cindex
				uf.arr[id] = id
				uf.sz[id] = 1
			}
		}
	}
	//union relevant elements
	for rindex, r := range board {
		for cindex, c := range r {
			if c == 'O' {
				id := rindex*col_num+cindex	
				//union super parent 
				if rindex == 0 || rindex == row_num - 1 || cindex == 0 || cindex == col_num - 1 {
					uf.union(super_parent, id)
				} 
				id, right, down := rindex*col_num+cindex, rindex*col_num+cindex+1, (rindex+1)*col_num+cindex
				if cindex + 1 < col_num && board[rindex][cindex+1] == 'O' {
					uf.union(id, right)
				}
				if rindex + 1 < row_num && board[rindex+1][cindex] ==  'O' {
						uf.union(id, down)
				}
			}
		}
	}
	//flip connected  part
	for rindex, r := range board {
		for cindex, c := range r {
			if c == 'O' {
				id := rindex*col_num+cindex	
				if rindex == 0 || rindex == row_num - 1 || cindex == 0 || cindex == col_num - 1 {
					continue
				} else {
					if !uf.connect(id, super_parent) {
						board[rindex][cindex] = 'X'
					}
				}
			}
		}
	}

}
