package tree
import "strconv"


func binaryTreePaths(root *TreeNode) []string {
	   res := []string{}
	   if root == nil {
		   return res
	   }	
	   val := strconv.Itoa(root.Val)
	   if  root != nil && root.Left == nil && root.Right == nil {
		   return []string{val}
	   }
	   if root.Left != nil {
		   for _, p := range binaryTreePaths(root.Left) {
			   res = append(res, val +"->"+p)
		   }
	   }
	   if root.Right != nil {
			for _, p := range binaryTreePaths(root.Right) {
				res = append(res, val +"->"+p)
			} 
	   }
	   return res
}