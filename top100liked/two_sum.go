package main

/**
   When  iterate the list, check whether a target - ele in HashMap:
	  True: you found the pair
	  False: you should store the  ele in the HashMap
*/
func twoSum(nums []int, target int) []int {
	finder := make(map[int]int)
	for i, value := range nums {
		out, ok := finder[target-value]
		if !ok {
			finder[value] = i
		} else {
			return []int{out, i}
		}
	}
	return nil
}
