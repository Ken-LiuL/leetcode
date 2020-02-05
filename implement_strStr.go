package str

func populateLPSArray(needle string, table []int) {
	len, table[0] = 0, 0
	i = 1
	for i < len(needle) {
		if needle[i] == needle[len] {
			len ++
			lps[i] = len
			i++
		} else {
			if len != 0 {
				//the logic of this line is:
				//for example, if needle is 
				//abaaba b
				//when b != a, then we know the maximum must smaller than 3
				//we need to find something like x1x2x0  == x3x4b
				//so x1x2 == x3x4 and x0 == b
				//x1x2 ==  x3x4  and aba == aba--> means it is the lps of aba
				//so we just check  needle[lps[len-1]] == b
				//then we know we have found x0 == b and x0 == needle[lps[len-1]]
				len = lps[len-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
}

//use KMP algorithm
//1. build prefix-postfix table 
//2. compare one by one and jump based on number in table
func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    if haystack == "" {
        return -1
    }
	// build prefix-postfix table first
	table := make([]int, len(needle))
	populateLPSArray(needle, table)

	i, j := 0, 0
	//iterate and compare
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			if len(needle) - 1 == j {
				return i-len(needle)+1
			} else {
				i, j = i+1, j+1
			}
		} else {
			if j == 0 {
				i++
			} else {
				i, j = i - table[j-1], 0
			}
		}
	}
    return -1
	
}