package UnionFind

type UnionFind struct {
	arr []int
	sz  []int
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
			uf.arr[rq] = rp
			uf.sz[rp] += uf.sz[rq]
		}
		uf.cnt--
	}
}
func (uf *UnionFind) connect(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

func findRedundantConnection(edges [][]int) []int {
	uf := NewUnionFind(len(edges) + 1)
	for i := 0; i < len(edges)+1; i++ {
		uf.arr[i] = i
		uf.sz[i] = 1
	}
	for _, e := range edges {
		if uf.connect(e[0], e[1]) {
			return e
		}
		uf.union(e[0], e[1])
	}
}
