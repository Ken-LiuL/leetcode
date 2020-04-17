package math

func min(a, b, c int) int{
	res := a
	if b < a {
		res = b
	}
	if c < res {
		res = c
	}
	return res
}

func nthUglyNumber(n int) int {
	res := make([]int, 0, n)
	p2, p3, p5 := 0, 0, 0
	res = append(res, 1) 
	for n > 1 {
	   v2, v3, v5 := res[p2]*2, res[p3]*3, res[p5]*5
	   ugly := min(v2, v3, v5)
	   if ugly == v2 {
		   p2++
	   } 
	   if ugly == v3 {
		   p3++
	   } 
	   if ugly == v5 {
		   p5++
	   }
	   res = append(res, ugly)
	   n--
	}
	return res[len(res) - 1]
}