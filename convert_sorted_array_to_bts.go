package tree

/**
*  Balanced BST, is that, for each node, the height difference would not be larger than 1
*/
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
	return sortHelper(nums, 0, len(nums)-1)
}

func sortHelper(nums []int, start, end int) *TreeNode{
    if start > end {
        return nil
    }
	pivot := (start+end)/ 2
	n := &TreeNode{Val: nums[pivot]} 
	n.Left =  sortHelper(nums, start, pivot-1)
	n.Right = sortHelper(nums, pivot+1, end) 
	return n
}