package sort

import "math/rand"

// 912. Sort an Array 排序数组
// 给你一个整数数组 nums，请你将该数组升序排列。
// 使用快速排序 , 不使用内置函数

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// ====================原生的快速排序============================

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	// 随机选择基准元素
	randomPivot := left + rand.Intn(right-left+1)
	// 将基准元素放在最右边 - 为什么?
	// 因为快速排序的基准元素选择有多种方式，这里选择最右边的元素作为基准元素
	nums[right], nums[randomPivot] = nums[randomPivot], nums[right]
	
	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	// 选择最右边的元素作为基准元素
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

// 时间复杂度：平均 O(nlogn)，最坏情况下仍为 O(n²)，但随机化后最坏情况出现概率极低
// 空间复杂度：O(logn)

// ====================三路快排 - Array.sort()============================

func quickSort3Way(nums []int, left, right int) {
	if left >= right {
		return
	}
	
	lt, gt := partition3Way(nums, left, right)
	quickSort3Way(nums, left, lt-1)
	quickSort3Way(nums, gt+1, right)
}

func partition3Way(nums []int, left, right int) (int, int) {
	pivot := nums[right]
	lt := left
	gt := right
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		} else if nums[j] > pivot {
			nums[j], nums[gt] = nums[gt], nums[j]
			gt--
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return lt, gt
}
