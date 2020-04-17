package UnionFind


type UnionFind struct {
	
	arr  []int
	cnt int
}

func NewUnionFind(size int) *UnionFind {
   uf := &UnionFind{arr: make([]int, size)}
 
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
		 uf.arr[rq]  = rp
		 uf.cnt -= 1
	 }
}
//using union find algorithm to find the result
func numIslands(grid [][]byte) int {
    if len(grid) == 0{
        return 0
    }
   rl, cl := len(grid), len(grid[0])
   uf := NewUnionFind(rl*cl)
   //init number of componenets
   for rindex, r := range grid {
	   for cindex, c := range r {
		   if c == '1' {
			   id := rindex*cl+cindex
			   uf.arr[id] =  id
			   uf.cnt++
		   }
	   }
   } 
   //iterate grid
   for rindex, r := range grid {
	   for cindex, c := range  r {
		   if c == '1' {
			   id, right, down := rindex*cl+cindex, rindex*cl+cindex+1, (rindex+1)*cl+cindex
			   if cindex + 1 < cl && grid[rindex][cindex+1] == '1' {
				   uf.union(id, right)
			   }
			   if rindex + 1 < rl && grid[rindex+1][cindex] ==  '1' {
				   uf.union(id, down)
			   }
		   }
	   }
   }
   return uf.cnt
}
