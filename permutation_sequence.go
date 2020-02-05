package array


func getPermutation(n int, k int) string {
	//initialize
	pool := make([]int, n)
	for i:=1; i <= n; i++ {
		pool[i-1] = i 
	}
	res := make([]int, 0, n)
	divident := rank(n-1)
	slot := k / divident
	remain := k % divident
	//ceiling
	if remain != 1 {
		slot++
		res = res.append(pool[slot-1])
		pool = append(pool[:slot-1], pool[slot:]...)
		divident = remain
		n-- 
	} else {
		
	}
	
}

func rank(n int) int {
	sum := 1
	if n > 0 {
		sum *= n
		n--
	}
	return sum
}