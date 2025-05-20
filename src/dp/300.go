package dp

// 300. 最长递增子序列
// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

// dp[i] 代表以nums[i]结尾的最长递增子序列的长度
// 状态转移方程：dp[i] = max(dp[i], dp[j]+1)
// 初始化：dp[i] = 1
// 返回：max(dp)

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		maxLength = max(maxLength, dp[i])
	}
	return maxLength
}

// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
