package string

func convert(s string, numRows int) string {
	res := make([]byte, len(s))
	if numRows == 1 {
		return s
	}
	l := 0
	for i := 0; i < numRows; i++ {
        for j, k := 0, i; k < len(s); j++ {
			res[l] = s[k]
			l++
			//first of last line of the result
			if (i == 0 || j % 2 == 0) 	&& i != numRows - 1 {
				k += 2 * (numRows - i - 1)
			} else {
				k += 2*i
			}
		}
	}
	return string(res)
}