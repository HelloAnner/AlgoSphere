package enum

// 2016. 增量元素之间的最大差值
// 枚举右，维护左

func maximumDifference(nums []int) int {
	min := nums[0]
	maxNum := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] > min {
			maxNum = max(maxNum, nums[i]-min)
		} else {
			min = nums[i]
		}
	}
	return maxNum
}

