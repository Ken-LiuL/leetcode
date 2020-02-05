package tree

//  like for tree [a1, a2, a3]
//     helper(a2, a1) --->
//                      helper(nil, a2) ---> return a2
//                      res := a2      
//                      a2.Left = nil, a2.Right = helper(nil, a1)   --> return a1
//                      a2.Right = a1  		
//                      return a2 
//     res :=  a2
//     a1.Left = nil
//     helper(a3, nil) -->
//                      helper(nil, a3) --->  return a3
//                      res := a3
//                      a3.Left = nil, a3.Right = helper(nil,  nil) --> return nil
//                      return a3
//     a1.Right = a3
//  -------------------------------------------------------------------------------
//  a2.right = a1
//  a1.Right = a3
//  a3.Right = nil
//  -------------------------------------------------------------------------------
// The key point for this algorithm is that:
// 			1. if you traverse down left,  you should let the left child point to the parent,
//                      if this is a left-left, then the returned res is actually useless, however, if it is right-left
//					    in a structure like parent -> right, we should return the leftmost element of this right branch to parent
//                      so this is why the returned res matters
//          2. if you traverse down right, you should let the right child point to the parent of its leftmost parent      
//				so this is why w need to keep passing next from top to the very down place, so that it could point to 

func helper(root *TreeNode, next *TreeNode) *TreeNode{
	if root  == nil {
		return next
	}
	res := helper(root.Left, root)
	root.Left = nil
	root.Right = helper(root.Right, next)
	return res
}

func increasingBST(root *TreeNode) *TreeNode {
	return helper(root)
}