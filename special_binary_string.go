package main
import "strings"
import "sort"
import "fmt"
/*   
 *  1. the special string must starts with 1 because, every prefix of the special string has at least as many 1's and 0's, so if it starts with 0, then the prefix [0] would not meet the condition
 *  2. the special string must has the form 1M0, cause, if it is 1M1 form, then, since the number
 * of 0 should be equal to the number of 1, then there will be prefix with more number of 0 
 * special string is like ()(( )), valid parenthesis
 */
func makeLargestSpecial(s string) string {
	flag, start := 0, 0
	res := make([]string, 0, 5)
	for ind, r := range s {
		if r == '1'  {
			//((
			flag++
		} else {
			//))
			flag--
		}
		if flag == 0 {
			//if means  end of pair of parenthesis
			res = append(res, "1"+makeLargestSpecial(s[start + 1: ind])+"0")
			//flag the start of the next pair of parenthesis
			start  = ind+1
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(res)))
	return strings.Join(res, "")
}
