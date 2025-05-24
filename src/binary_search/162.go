package binary_search

// 162. 寻找峰值
// 峰值元素是指其值严格大于左右相邻值的元素。


func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// 时间复杂度：O(log n)
// 空间复杂度：O(1)
