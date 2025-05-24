package kadane

// 53. 最大子数组和
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。


func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := 0

	for _, num := range nums {
		if currentSum < 0 {
			currentSum = num
		} else {
			currentSum += num
		}
		
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
