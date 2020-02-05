package recursion

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Rgiht *TreeNode
}

var memorization = make([][]*TreeNode, 20)

/**
*  The keypoint is for any N, the combination is (Left(1), Right(N - 1)), (Left(3), Right(N-3)), ...
*  when N == 1, it means there is no children under this node, like for instance Left(1)
 */
func allPossibleFBT(N int) []*TreeNode {
	if memorization[N] != nil {
		return memorization[N]
	}
	N -= 1
	if N == 0 {
		return []*TreeNode{&TreeNode{}}
	}

	res := make([]*TreeNode, 0, 10)
	for i := 1; i < N; i += 2 {
		if memorization[i] == nil {
			memorization[i] = allPossibleFBT(i)
		}
		if memorization[N-i] == nil {
			memorization[N-i] = allPossibleFBT(N - i)
		}
		for _, left := range memorization[i] {
			for _, right := range memorization[N-i] {
				root := &TreeNode{}
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}
	}
	return res
}
