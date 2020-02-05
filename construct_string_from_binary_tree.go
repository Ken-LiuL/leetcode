package tree

import "strconv"
import "strings"

var res []string
func preorder(t *TreeNode) *TreeNode {
	if t == nil {
		return t
	}
	res = append(res, strconv.Itoa(t.Val))
	res = append(res, "(")
	lr := preorder(t.Left)
	res = append(res, ")")
	res = append(res, "(")
	rr := preorder(t.Right)
	res = append(res, ")")
	if  lr == nil && rr == nil {
		res = res[:len(res) - 4]
	} else if rr == nil {
		res = res[:len(res)-2]
	}
	return t
}

func tree2str(t *TreeNode) string {
	res = make([]string, 0, 20)
	preorder(t)
	return strings.Join(res, "")
}