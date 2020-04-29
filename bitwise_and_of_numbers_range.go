package bit

const powers = int{1073741824,536870912,268435456,134217728,67108864,33554432,16777216,8388608,4194304,2097152,1048576,524288,262144,131072,65536,32768,16384,8192,4096,2048,1024,512,256,128,64,32,16,8,4,2,1,0}
func rangeBitwiseAnd(m int, n int) int {
	res := 0
	val := powerOf2(31)
	pre := 2147483648
	for i := 31; i >= 1; i-- {
		val = powers[31-i]
		if val <= m && n < pre {
			res += val
			m -= val
			n -= val
		}
		pre = val
	}
	return res
}
//obervation:
// m could be  xxx0...
// n could be  xxx1...
// they should share a common prefix bits

func rangeBitwiseAnd2(m int, n int) int {
	step := 0
	 for m != n {
		 m >>=1
		 n >>=1
		 step++
	 }
	 return m << step
}


//we could use Brian Kernighan’s Algorithm: to count bits
func rangeBitwiseAnd3(m int, n int) int {
	for n > m {
		n &= (n - 1)
	}
	return m & n
}
