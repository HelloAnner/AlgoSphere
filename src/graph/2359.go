package graph

// 2359. 找到离给定两个节点最近的节点
// 给你一个 n 个节点的 有向图 ，节点编号为 0 到 n - 1 ，每个节点 至多 有一条出边。
// 有向图用大小为 n 下标从 0 开始的数组 edges 表示，表示节点 i 有一条有向边指向 edges[i] 。如果节点 i 没有出边，那么 edges[i] == -1 。
// 返回一个从 node1 和 node2 都能到达节点的编号，使节点 node1 和节点 node2 到这个节点的距离 较大值最小化
// 如果答案不止一个，请返回节点 编号最小的 那个。
// 两个节点之间可能有多条路径，对于每一个节点，至少有一条路径从该节点出发且通向一个 有效 终点。

// 较大值最小化:
// 1. 从node1出发，找到所有能到达的节点
// 2. 从node2出发，找到所有能到达的节点
// 3. 找到两个集合的交集，并返回交集中节点编号最小的那个
// 4. 如果两个集合没有交集，返回-1

func closestMeetingNode(edges []int, node1 int, node2 int) int {
    n := len(edges)
    dist1 := make([]int, n)
    dist2 := make([]int, n)
    
    for i := range dist1 {
        dist1[i] = -1
        dist2[i] = -1
    }
    
    // 使用BFS计算从node1出发的最短距离
    queue := []int{node1}
    dist1[node1] = 0
    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:]
        v := edges[u]
        if v != -1 && dist1[v] == -1 {
            dist1[v] = dist1[u] + 1
            queue = append(queue, v)
        }
    }
    
    // 使用BFS计算从node2出发的最短距离
    queue = []int{node2}
    dist2[node2] = 0
    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:]
        v := edges[u]
        if v != -1 && dist2[v] == -1 {
            dist2[v] = dist2[u] + 1
            queue = append(queue, v)
        }
    }
    
    ans := -1
    minDist := n + 1
    for i := 0; i < n; i++ {
        if dist1[i] != -1 && dist2[i] != -1 {
            currentMax := max(dist1[i], dist2[i])
            if currentMax < minDist || (currentMax == minDist && i < ans) {
                minDist = currentMax
                ans = i
            }
        }
    }
    
    return ans
}