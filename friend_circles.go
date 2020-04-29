package UnionFind


type UnionFind struct {
	arr  []int
	sz   []int
	cnt int
}

func NewUnionFind(size int) *UnionFind {
   uf := &UnionFind{arr: make([]int, size), sz: make([]int, size)}
   uf.cnt = size
   return uf
}

func (uf *UnionFind) root(p int) int {
	for p != uf.arr[p] {
		uf.arr[p] = uf.arr[uf.arr[p]]
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
		 uf.cnt--
	}
}
func (uf *UnionFind) connect(p, q int) bool {
	return uf.root(p) == uf.root(q)
}


func findCircleNum(M [][]int) int {
	row_num, col_num := len(M), len(M[0])
	uf := NewUnionFind(row_num)
	for i := 0; i < row_num; i++ {
		uf.arr[i] = i
	}
	for i := 0; i < row_num; i++ {
		for j := i; j < col_num; j++ {
			if M[i][j] == 1 {
			 uf.union(i, j)
			}
		}
	}
	return uf.cnt

}


