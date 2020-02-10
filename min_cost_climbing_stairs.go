package array

var memorization map[int]int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	memorization = make(map[int]int)
	return helper(cost)
}

func helper(cost []int) int {
	if len(cost) <= 1 {
		return 0
	}
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	if _, ok := memorization[len(cost)]; !ok {
		memorization[len(cost)] =  min(cost[0] + helper(cost[1:]), cost[1] + helper(cost[2:]))
	}
	return memorization[len(cost)]
}