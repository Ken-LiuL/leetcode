package graph

const MaxInt = int(^uint(0) >> 1)
func minDistance(costs []int, visited []int, N int) int {
	m, mInd := MaxInt, -1
	for i := 0; i < N; i++ {
		if  costs[i] < m && visited[i] == 0 {
			m = costs[i]
			mInd = i
		}
	}
	return mInd
}
//Dijkstra's algorithm
func networkDelayTime(times [][]int, N int, K int) int {
	//init the graph with dict{n -> {n: cost}}
	 graph := make(map[int]map[int]int, N)
	 for _, n := range times {
		 if _, ok := graph[n[0]]; ok {
			graph[n[0]][n[1]] = n[2]
		 } else {
		   graph[n[0]] =  make(map[int]int, N)
		   graph[n[0]][n[1]] = n[2]
		 }
	 }
	 costs := make([]int, N)
	 for i := 0; i < len(costs); i++ {
		 costs[i] =  MaxInt
	 }
	 visited := make([]int,  N)
	 //set cost and visited
	 costs[K-1] = 0
     for i := 0; i < N; i++ {
		  chosen := minDistance(costs, visited, N)
          if chosen == -1 {
              break
          }
		  visited[chosen] = 1
		  for child, cos  := range graph[chosen+1] {
              child := child-1
			  if visited[child] == 0 && costs[child] > costs[chosen] + cos {
				  costs[child] = costs[chosen] + cos
			  }
		  }
	 }
    res := costs[0]
    for _, c := range costs[1:]{
        if c > res {
            res = c
        }
    }
    if res == MaxInt {
        return -1
    }
    return res
}