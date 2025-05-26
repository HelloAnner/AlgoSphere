package graph

// 130. 被围绕的区域
// 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

func solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				dfs1(board, i, j)
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs1(board [][]byte, i, j int) {
	rows := len(board)
	cols := len(board[0])

	if i < 0 || i >= rows || j < 0 || j >= cols || board[i][j] == 'X' {
		return
	}

	board[i][j] = 'X'
	dfs1(board, i+1, j)
	dfs1(board, i-1, j)
	dfs1(board, i, j+1)
	dfs1(board, i, j-1)
}
