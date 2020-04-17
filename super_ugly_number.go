package array
func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}


func nthSuperUglyNumber(n int, primes []int) int {
    res := make([]int, n)
	pointers := make([]int, len(primes))
    res = append(res, 1) 
	for n > 1 {
	   minValue := primes[0]*res[pointers[0]]
	   for i := 1; i < len(primes); i++ {
			minValue = min(minValue, primes[i]*res[pointers[i]])
	   }
	   for i := 0; i < len(primes);i++ {
		   if minValue == primes[i] * res[pointers[i]] {
			  pointers[i]++
		   }
		}
		
	   res = append(res, minValue)
	   n--
	}
	return res[len(res) - 1]
}