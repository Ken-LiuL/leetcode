package array

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, n := range A {
		sum += n
	}
	sum /= 3
	start, end := 0, len(A) -1
	first, last := 0, 0
	for start < end {
		if first != sum  {
			first += A[start]
			start++
		}
		if last != sum {
		   last += A[end]
		   end-- 
		}
		if first == sum && last == sum {
			return true
		} 
	}
	return false
}