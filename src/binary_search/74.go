package binary_search

// 74. 搜索二维矩阵
// 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
// 每行中的整数从左到右按升序排列。

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		left, right := 0, cols-1
		for left <= right {
			mid := (left + right) / 2
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

// 时间复杂度：O(m * log n)
// 空间复杂度：O(1)
