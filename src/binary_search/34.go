package binary_search

// 34. 在排序数组中查找元素的第一个和最后一个位置
// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。


func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			left = mid
			right = mid
			for left >= 0 && nums[left] == target {
				left--
			}
			for right < len(nums) && nums[right] == target {
				right++
			}
			return []int{left + 1, right - 1}
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{-1, -1	}
}

// 时间复杂度：O(log n)
// 空间复杂度：O(1)
