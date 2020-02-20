package array


func duplicateZeros(arr []int)  {
	head := len(arr)-1
	tail := len(arr) - 1

	for i := 0; i <= head; i++ {
		 if arr[i] == 0  {
			if i != head {
			  head--
			} else {
			  //handle corner case	
			  arr[tail] = arr[head]
			  tail--
			  head--
			}
		 }
	 } 
	 for head < tail && head >= 0 {
		 if arr[head] == 0  { 
			 arr[tail] = 0
			 tail--
			 arr[tail] = 0
		 } else { 
		     arr[tail] = arr[head]
     		 tail--
			 head--
		 }
	 }
}