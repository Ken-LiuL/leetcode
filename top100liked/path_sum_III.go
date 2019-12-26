package tree



func helper(root *TreeNode, sum int, target int, cache map[int]int) int{
	if root == nil {
		return 0
	}
	//sum from  root to current node
	sum += root.Val
	res := 0
	//     sum - target               target 
	//root---------------->some node----------> currentNode
	//--------- ------------sum ----------------------------
	//so we try to find whether there are some node, so that some node to current node is target
	if val, ok := cache[sum - target]; ok {
		res = val
	}
	//update cache
	if val, ok := cache[sum]; ok {
		cache[sum] = val + 1
	} else {
		cache[sum] = 1
	}
	//recursively
	res +=  helper(root.Left, sum, target, cache) + helper(root.Right, sum, target, cache)
	//update cache
	cache[sum] = cache[sum] -  1
	return res


}

//1. The prefix stores the sum from the root to the current node in the recursion
//2. The map stores <prefix sum, frequency> pairs before getting to the current node. We can imagine a path from the root to the current node. The sum from any node in the middle of the path to the current node = the difference between the sum from the root to the current node and the prefix sum of the node in the middle.
//3. We are looking for some consecutive nodes that sum up to the given target value, which means the difference discussed in 2. should equal to the target value. In addition, we need to know how many differences are equal to the target value.
//4. Here comes the map. The map stores the frequency of all possible sum in the path to the current node. If the difference between the current sum and the target value exists in the map, there must exist a node in the middle of the path, such that from this node to the current node, the sum is equal to the target value.
//5. Note that there might be multiple nodes in the middle that satisfy what is discussed in 4. The frequency in the map is used to help with this.
//6. Therefore, in each recursion, the map stores all information we need to calculate the number of ranges that sum up to target. Note that each range starts from a middle node, ended by the current node.
//7. To get the total number of path count, we add up the number of valid paths ended by EACH node in the tree.
//8. Each recursion returns the total count of valid paths in the subtree rooted at the current node. And this sum can be divid
func pathSum(root *TreeNode, sum int) int {
	 cache := make(map[int]int, 100)
	 cache[0] = 1
	 return helper(root, 0, sum, cache)
}