package backtrace

// 79. 单词搜索
// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。


func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrace5(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func backtrace5(board [][]byte, word string, i int, j int, index int) bool {
	if index == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[index] {
		return false
	}

	board[i][j] = '#'

	found := backtrace5(board, word, i+1, j, index+1) ||
		backtrace5(board, word, i-1, j, index+1) ||
		backtrace5(board, word, i, j+1, index+1) ||
		backtrace5(board, word, i, j-1, index+1)

	board[i][j] = word[index]

	return found
}



