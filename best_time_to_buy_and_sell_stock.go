package array

//we could also use Kadane's Algorithm
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
        return 0
    }
	start,  bestProfit := prices[0], 0
	for _,  p := range prices[1:] {
		 if start > p {
			 start = p
		 } else {
			 bestProfit = max(p - start, bestProfit)
		 }
	}
	return bestProfit
}