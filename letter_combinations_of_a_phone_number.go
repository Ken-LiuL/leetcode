package backtracking


var keys []string = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	res := make([]string)
	helper("", digits, &res)
	return res
}

func helper(prefix string, digits string, res *[]string) {
	if len(digits) == 0 {
		*res = append(*res, prefix)
	}
	letters := keys[digits[0] - '0']
	for i := 0; i < len(letters);i++ {
		helper(string(letters[i])+prefix, digits[1:], res)
	}
}