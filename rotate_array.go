package array


func forward(nums []int, step int) {
	if step == 0 {
		return 
	}
	tmp := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i-1] = nums[i]
	}
	nums[len(nums) - 1] = tmp
	forward(nums, step-1)
}

func backward(nums []int, step int) {
	if step == 0 {
		return
	}
	tmp := nums[len(nums) - 1]
	for i := len(nums) - 1; i > 0; i -- {
		nums[i] = nums[i-1]
	}
	nums[0] = tmp
	backward(nums, step - 1)
}

func rotate(nums []int, k int)  {
	l := len(nums)
	//get the actual step 
	k = k % l
	//check 
	if k > l/2 {
    //move forward
		forward(nums, l - k)
	} else {
    //move backward
		backward(nums, k)
	}
}