package tree


const MaxInt = int(^uint(0) >> 1)

func collect(node *TreeNode, arr *[]int ){
	if node == nil {
		return
	}
	collect(node.Left, arr)
	*arr = append(*arr, node.Val)
	collect(node.Right, arr)
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1, arr2 := make([]int, 0, 20), make([]int, 0, 20)
	collect(root1, &arr1)
	collect(root2, &arr2)

	res := make([]int, 0, len(arr1)+len(arr2))
	ptr1, ptr2 := 0, 0
	//then we need to merge this two array
	for ptr1 < len(arr1) || ptr2 < len(arr2) {
		 n1, n2 := MaxInt, MaxInt
		 if ptr1 < len(arr1) {
			 n1 =  arr1[ptr1]
		 }
		 if ptr2 < len(arr2) {
			 n2 = arr2[ptr2]
		 }
		 if n1 < n2 {
			res = append(res, n1)
			ptr1++
		 } else {
			res = append(res, n2)
			ptr2++
		 }	
	}
	return res
}