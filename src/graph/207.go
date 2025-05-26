package graph

// 207. 课程表
// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，
// 其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程 bi 。

func canFinish(numCourses int, prerequisites [][]int) bool {
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

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visited[node] = true
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			return false
		}
	}

	return true
}

// 时间复杂度：O(n + m)
// 空间复杂度：O(n + m)

