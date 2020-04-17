package graph


const MaxInt = int(^uint(0) >> 1)

//BFS algorithm
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	  res := MaxInt 
	  graph := make(map[int]map[int]int, 100)
	  //the graph
	  for _, trip := range flights {
		  if _ , ok := graph[trip[0]]; ! ok {
			graph[trip[0]] = make(map[int]int, 10)
		  }
		  graph[trip[0]][trip[1]] = trip[2]
	  }
	  queue := make([][]int, 0, n)
	  queue = append(queue, []int{src, 0})
	  step := 0
	  for len(queue) != 0 {
		  size := len(queue)
		  for i := 0; i < size; i++ {
			  cur := queue[i]
			  if cur[0] == dst && cur[1] < res {
				  res = cur[1]
			  }
			  for child, cost := range graph[cur[0]] {
				  if cost +  cur[1] <= res {
				      queue = append(queue, []int{child, cur[1] + cost})
				  }
			  }
		  }
		  if step > K {
			  break
		  }
		  step += 1
		  queue = queue[size:]
	  }
	  if res == MaxInt {
		  return -1
	  }
	  return res
}

//bellman ford algorithm
//bellman ford algorithm
func findCheapestPrice2(n int, flights [][]int, src int, dst int, K int) int {
	costs := make([]int, n)
	for i := 0; i < len(costs); i++ {
		costs[i] = MaxInt
	}
	costs[src] = 0

	for i := 0; i <= K; i++ {
        newCost := make([]int, n)
		copy(newCost, costs)
		for _, edge := range flights { 
			if costs[edge[0]] + edge[2] < newCost[edge[1]] {
				newCost[edge[1]] = costs[edge[0]] + edge[2]
			}
		}
		costs = newCost
	}
	if costs[dst] == MaxInt {
		return -1
	}
	return  costs[dst]
}