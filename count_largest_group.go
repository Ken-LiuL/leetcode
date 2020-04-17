package array

func sumOfDigit(a int) int {
	 s := 0
	 for a > 0 {
		 s += a % 10
		 a = a / 10
	 }
	 return s
}

func countLargestGroup(n int) int {
	cache := make(map[int]int)
	largest := 0

	for i := 1; i <= n; i++ {
		s := sumOfDigit(i)
		cache[s] += 1
		if cache[s] > largest {
			largest = cache[s]
		}
	}
	cnt := 0
	for  _, v := range cache {
		if v  == largest {
			cnt += 1
		}
	}
	return cnt

}