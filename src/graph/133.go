package graph

// 133. 克隆图
// 给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。
// 图中的每个节点都包含它的值 val（int） 和其邻居列表（list[Node]）。

type Node struct {
	Val int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	return dfs3(node, visited)
}

func dfs3(node *Node, visited map[*Node]*Node) *Node {
	if visited[node] != nil {
		return visited[node]
	}

	cloneNode := &Node{
		Val: node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = cloneNode

	for i, neighbor := range node.Neighbors {
		cloneNode.Neighbors[i] = dfs3(neighbor, visited)
	}

	return cloneNode

}
