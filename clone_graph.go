package graph


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

 //BFS
 func cloneGraph(node *Node) *Node {
    if node == nil {
        return node
    }
	 visited := make(map[*Node]*Node, 100)
	 queue := make([]*Node, 0, 100)
	 queue = append(queue, node)
     copyRoot := &Node{Val: node.Val, Neighbors: make([]*Node, 0, 10)}
	 visited[node] =  copyRoot
	 for len(queue) != 0 {
		 n := queue[0]
		 for _, neighbor := range n.Neighbors {
			 if copyN, ok := visited[neighbor]; !ok {
                 visited[neighbor] = &Node{Val: neighbor.Val, Neighbors:make([]*Node, 0, 10)}
                 visited[n].Neighbors = append(visited[n].Neighbors, visited[neighbor])
			     queue = append(queue, neighbor) 
			 } else {
				visited[n].Neighbors = append(visited[n].Neighbors, copyN)
			 }
		 }
		 queue = queue[1:]
	 }
	 return copyRoot
}