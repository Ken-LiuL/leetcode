package string


func scoreOfParentheses(S string) int {
	counter := 0
	cnt := 0
    for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			counter++
		} else {
            counter--
			if S[i - 1] == '(' {
				cnt += 1 << counter
			}
		}
	}
    return cnt
}