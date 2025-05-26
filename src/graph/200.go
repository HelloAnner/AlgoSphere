package graph

// 200. 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	rows := len(grid)
	cols := len(grid[0])

	if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

// 时间复杂度：O(n * m)
// 空间复杂度：O(n * m)
