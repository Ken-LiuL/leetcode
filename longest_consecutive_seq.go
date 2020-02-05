package array


//the keypoint is about inserting number to slot
//for examples, 100 4 200 1 3 2
//counts[100] = 1, represent one
//counts[4] = 1, represent one
//counts[200] = 1, represent one
//when 3 show up, it should connect  to 4 to form 3 4
//so  counts[3+1] is equal to 1, counts[3]  = 1 + 1 as 2, represent the length of 3 4
//and we should update counts[4] to 2
//finally, when 2 show up, it should connect  3 4 and 1 to form  1 2 3 4
//since counts[3] is equal to 2, counts[1] is equal to 1, so counts[2] is equal to 4
//and then update counts[2 - 1] as 4, and counts[2 + 2] as 4, so that, we have
//counts[100] = 1
//counts[1] = 4
//counts[2] = 4
//counts[3] = 2
//counts[4] = 4
//key point, let boundary number always hold the length of the sequence
func longestConsecutive(nums []int) int {
	counts := make(map[int]int, len(nums))
	max := 0
	for _, n := range nums {
		if _, ok := counts[n]; ok {
			continue
		}
		down := 0
		up := 0
		if d, ok := counts[n - 1]; ok {
			down = d
		}
		if u, ok := counts[n + 1]; ok {
			up = u
		}
		count := up + down +  1
		counts[n] = count
		counts[n + up] = count
		counts[n - down] = count
		if count > max {
			max = count
		}
	}
	return max
	 
}