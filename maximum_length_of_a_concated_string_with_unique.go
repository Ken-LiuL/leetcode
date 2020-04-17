package bit



func bitCount(a int) int {
	if a == 0 {
		return 0
	}
	return (a & 1) + bitCount(a >> 1)
}

func maxLength(arr []string) int {
	 bits := make([]int, 0, len(arr))
	 for _, str := range arr {
		   bit := 0 
		   uniq := true
		   for _, c := range str {
			   if  (bit >> (c - 'a') & 1) == 1 {
				   uniq = false
				   break
			   }
			   bit = (1 << (c-'a')) | bit
		   }
		   if uniq {
		   	 bits = append(bits, bit)
		   }
	 } 
	 res := helper(bits)
	 m := 0
	 for _, r := range res {
		if bitCount(r) > m {
			m = bitCount(r)
		}
	 }
	 return m
}

func helper(arr []int) []int {
    if len(arr) == 0 {
        return []int{}
    }
   if len(arr) == 1 {
	   return []int{arr[0]}
   }
   res:= helper(arr[1:])
   l := len(res)
   res = append(res, arr[0])
   for i := 0; i < l; i++ {
	   if res[i] & arr[0] == 0 {
		   res = append(res, res[i] | arr[0])
	   }    
	}
   return res
}