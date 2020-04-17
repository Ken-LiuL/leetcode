package UnionFind

type UnionFind struct {
	arr  []int
	sz   []int
	cnt int
}

func NewUnionFind(size int) *UnionFind {
   uf := &UnionFind{arr: make([]int, size), sz: make([]int, size)}
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
func (uf *UnionFind) getComponents() int {
	return uf.cnt
}

func (uf *UnionFind) connect(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

//https://buptwc.com/2018/11/25/Leetcode-947-Most-Stones-Removed-with-Same-Row-or-Column/
func removeStones(stones [][]int) int {
   //number of stone - (nummber of isolated island)  
   uf := NewUnionFind(len(stones))
   for i, stone := range stones {
	   uf.arr[i] = i
	   uf.cnt++
   }
   for i, s1 := range stones {
	   for j, s2 := range stones {
		   if s1[0] == s2[0] || s1[1] == s2[1] {
			   uf.union(i, j)
		   }
	   }
   }
   
   return len(stones) - uf.getComponents()
}
