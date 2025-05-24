package bit

// 137. 只出现一次的数字 II
// 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

// 第一次遇到某位为1：ones 记录该位，twos 保持0。
// 第二次遇到某位为1：twos 记录该位，ones 清零。
// 第三次遇到某位为1：ones 和 twos 均清零。

func singleNumberII(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = (ones ^ num) & ^twos
		twos = (twos ^ num) & ^ones
	}
	return ones
}
