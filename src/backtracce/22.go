package backtrace

// 22. 括号生成
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。


func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	ans := []string{}
	backtrace4(n, 0, 0, "", &ans)
	return ans
}

func backtrace4(n int, left int, right int, path string, ans *[]string) {
	if len(path) == n*2 {
		*ans = append(*ans, path)
		return
	}

	if left < n {
		path += "("
		backtrace4(n, left+1, right, path, ans)
		path = path[:len(path)-1]
	}

	if right < left {
		path += ")"
		backtrace4(n, left, right+1, path, ans)
		path = path[:len(path)-1]
	}
}
