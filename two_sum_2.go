package array


func twoSum(numbers []int, target int) []int {
	start, end := 0, len(numbers) - 1
	for start < end {
		res := numbers[start] + numbers[end] 
		if res > target {
			end--
		} else if res < target {
			start++
		} else {
			return []int{start+1, end+1}
		}
	}
	return nil
}