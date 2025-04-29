package window

import "math"

// https://leetcode.cn/problems/minimum-size-subarray-sum/?envType=study-plan-v2&envId=top-interview-150
// 给定一个含有 n 个正整数的数组和一个正整数 target 。
// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。
// 如果不存在符合条件的子数组，返回 0 。
// 思路: 滑动窗口
// 1. 使用两个指针 left 和 right 表示当前窗口的左右边界
// 2. 使用一个变量 sum 来记录当前窗口内元素的和
// 3. 移动 right 指针，扩大窗口，更新 sum
// 4. 当 sum 大于等于 target 时，移动 left 指针，缩小窗口，更新 sum
// 5. 记录最小的窗口长度

func MinSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	minLen := math.MaxInt32

	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}

	if minLen == math.MaxInt32 {
		return 0
	}

	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
