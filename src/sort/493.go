package sort

// 493. 翻转对
// 给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。
// 你需要返回给定数组中重要翻转对的数量。
// 给定数组的长度不会超过50000。
// 暴力计算会超时
// 使用归并排序来计算


func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := left + (right-left)/2
	count := mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right)
	count += merge(nums, left, mid, right)
	return count
}

func merge(nums []int, left, mid, right int) int {
	// 计算重要翻转对的数量
	count := 0
	j := mid + 1
	for i := left; i <= mid; i++ {
		for j <= right && nums[i] > 2*nums[j] {
			j++
		}
		count += j - (mid + 1)
	}
	// 合并两个有序数组
	temp := make([]int, right-left+1)
	i, k := left, 0
	for j := mid + 1; j <= right; j++ {
		for i <= mid && nums[i] <= nums[j] {
			temp[k] = nums[i]
			k++
			i++
		}
		temp[k] = nums[j]
		k++
	}
	for i <= mid {
		temp[k] = nums[i]
		k++
		i++
	}
	for p := 0; p < len(temp); p++ {
		nums[left+p] = temp[p]
	}
	return count
}
