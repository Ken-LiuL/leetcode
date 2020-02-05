package main

//for an array, if all elements appear p times, but one appear k time,
//if p is even and k is odd, we could always use xor to solve,
//but for other condition, it becomes different

//we add the number of each bit of the number, and % p, if that bit finally is
//not zero, then we know this bit of the special number is 1 otherwise, it could
//be zero
//also consider this complex solution 
//
func singleNumber(nums []int) int {
	ans := 0 
	for i := 0; i < 32; i++ {
		sum := 0
		for _, n := range nums {
			if   ((n >> i) & 1) == 1 {
				sum++
				sum %= 3
			}
			//we do not need to consider the zero bit
		}
		if sum != 0 {
			ans |= sum << i
		}
	}
	return ans
}
//another better one
//https://leetcode.com/problems/single-number-ii/discuss/43294/Challenge-me-thx
//logic is like this:
//we would like use ones and two to form a two bit number as state machine
//and the state machine should be like 00 -> 01 -> 10 -> 00, when met 1, otherwise
//stay as same
//condition 1:
//     make sure, if either is one, the other must be zero, & ^other could make sure
//condition 2:
//     make sure, if either met zero, it keep the same, 
//      ones ^ 0 = ones, twos ^ 0 = twos
//    if ones is zero, then ones & ^twos must be zero
//    if ones is one, the twos must be one, so that, ones & ^twos is one
//    if twos is zero, then twos & ^ones must be zero
//    if twos is one, then ones must be zero, so twos & ^ones is one
func singleNumber(nums []int) int {
	ones , twos := 0, 0
	for _, n := range nums {
		ones = (ones ^ n) & ^twos
		twos = (twos ^ n) & ^ones
	}
	return ones
}