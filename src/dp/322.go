package dp

// 322. 零钱兑换
// 给你一个整数数组 coins 表示不同面额的硬币，以及一个整数 amount 表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

// dp[i] 代表凑成总金额i所需的最少的硬币个数
// 状态转移方程：dp[i] = min(dp[i], dp[i-coin]+1)
// 初始化：dp[0] = 0
// 返回：dp[amount]

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// 时间复杂度：O(n*amount)
// 空间复杂度：O(amount)
