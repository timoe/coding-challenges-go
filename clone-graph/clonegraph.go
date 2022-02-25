package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	if len(node.Neighbors) == 0 {
		return &Node{Val: node.Val, Neighbors: []*Node{}}
	}

	visited := make(map[*Node]bool)
	references := make(map[int]*Node)

	clone := &Node{Val: node.Val}
	references[node.Val] = clone

	doClone(node, clone, visited, references)
	return clone
}

func doClone(node, target *Node, visited map[*Node]bool, references map[int]*Node) {
	_, ok := visited[node]

	if !ok {
		visited[node] = true
		for _, n := range node.Neighbors {
			node, ok := references[n.Val]
			cloned := node
			if !ok {
				cloned = &Node{Val: n.Val}
				references[n.Val] = cloned
			}
			target.Neighbors = append(target.Neighbors, cloned)
			doClone(n, cloned, visited, references)
		}
	}
}
