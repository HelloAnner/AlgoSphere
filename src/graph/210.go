package graph

// 210. 课程表 II
// 现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。
// 给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 构建邻接表
	graph := make([][]int, numCourses)
	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
	}

	// 记录每个节点的入度
	inDegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		inDegree[pre[0]]++
	}

	// 使用队列进行拓扑排序
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	visited := make([]bool, numCourses)
	result := make([]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visited[node] = true
		result = append(result, node)
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			return []int{}
		}
	}

	return result
}

// 时间复杂度：O(n + m)
// 空间复杂度：O(n + m)
