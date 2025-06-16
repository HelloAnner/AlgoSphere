package greedy

// 329. 矩阵中的最长递增路径
// 给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。


func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	maxPath := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			maxPath = max(maxPath, dfs(matrix, i, j, dp))
		}
	}

	return maxPath
}

// 深度优先搜索
// 记忆化搜索
// 从每个点出发，向四个方向进行深度优先搜索，记录每个点所能到达的最长递增路径长度
// 使用dp数组记录每个点所能到达的最长递增路径长度
// 如果dp[i][j] != 0，则直接返回dp[i][j]，否则进行深度优先搜索
// 深度优先搜索过程中，记录每个点所能到达的最长递增路径长度
// 返回dp[i][j]
func dfs(matrix [][]int, i, j int, dp [][]int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	rows, cols := len(matrix), len(matrix[0])
	dp[i][j] = 1
	for _, dir := range dirs {
		ni, nj := i + dir[0], j + dir[1]
		if ni >= 0 && ni < rows && nj >= 0 && nj < cols && matrix[ni][nj] > matrix[i][j] {
			dp[i][j] = max(dp[i][j], dfs(matrix, ni, nj, dp) + 1)
		}
	}

	return dp[i][j]
}

// 时间复杂度：O(m * n)
// 空间复杂度：O(m * n)