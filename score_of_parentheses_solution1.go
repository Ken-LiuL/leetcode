package string

func scoreOfParentheses(S string) int {
	counter := 0
	cnt := 0
	i := 0
	for i < len(S) {
		if S[i] == '(' {
			counter++
			i++
		} else {
			counter--
            cnt += pow(2, counter)
			i++
			for i < len(S) {
				if S[i] == ')' {
					counter--
                    i++
				} else {
					break
				}
			}
		}
	}
    return cnt
}