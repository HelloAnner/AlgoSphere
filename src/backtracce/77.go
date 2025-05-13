package backtrace

// 77. 组合
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。


func combine(n int, k int) [][]int {
	if k > n {
		return [][]int{}
	}
	ans := [][]int{}
	backtrace3(n, k, 1, []int{}, &ans)
	return ans
}

func backtrace3(n int, k int, start int, path []int, ans *[][]int) {
	if len(path) == k {
		*ans = append(*ans, append([]int{}, path...))
		return
	}

	for i := start; i <= n; i++ {
		path = append(path, i)
		backtrace3(n, k, i+1, path, ans)
		path = path[:len(path)-1]
	}
}
