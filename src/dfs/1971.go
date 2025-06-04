package dfs

// 1971. 寻找图中是否存在路径
// 有一个具有 n 个顶点的 双向 图。给你一个二维整数数组 edges ，其中 edges[i] = [ui, vi] 表示顶点 ui 和顶点 vi 之间存在一条 双向 边。
// 每个顶点都标记为 0 到 n - 1（包含 0 和 n - 1）。顶点 i 的值 相邻节点 是满足存在一条边从节点 i 到那个顶点的所有顶点。
// 注意顶点 不 包含它自己。
// 请你返回一个长度为 n 的数组 answer ，其中 answer[i] 是顶点 i 的值相邻节点的数目。

func validPath(n int, edges [][]int, source int, destination int) bool {
	visited := make([]bool, n)
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return dfs1(graph, visited, source, destination)
}

func dfs1(graph [][]int, visited []bool, source int, destination int) bool {
	if source == destination {
		return true
	}
	visited[source] = true
	for _, neighbor := range graph[source] {
		if !visited[neighbor] && dfs1(graph, visited, neighbor, destination) {
			return true
		}
	}
	return false
}