package dp

// 70. 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// dp[i] 代表爬到第i阶楼梯的方法数
// 状态转移方程：dp[i] = dp[i-1] + dp[i-2]
// 初始化：dp[0] = 1, dp[1] = 1
// 返回：dp[n]

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 时间复杂度：O(n)
// 空间复杂度：O(n)
