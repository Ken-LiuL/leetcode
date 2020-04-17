package tree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 //BFS
 type  N struct  {
	n *Node
	depth int
}

func maxDepth(root *Node) int {
   if root == nil {
	   return 0
   }
	queue := make([]*N, 0, 100)
	md := 0
	rn := &N{n: root, depth: 1}
	queue = append(queue, rn)
	for len(queue) != 0 {
		 node := queue[0]
		 if node.depth > md {
			 md = node.depth
		 }
		 for _, c := range node.n.Children {
			 queue = append(queue, &N{n: c, depth: node.depth+1})
		 }
		queue = queue[1:]
	}
	return md
}