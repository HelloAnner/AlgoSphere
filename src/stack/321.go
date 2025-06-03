package stack

// 321. 拼接最大数
// 给定长度分别为 m 和 n 的两个数组，其元素的取值范围分别是 [0,9] 和 [0,9]，
// 如果把两个数组拼接在一起，是否存在一种拼接方案，使得得到的数组最大。
// k 是两个数组拼接后的长度，k 的取值范围是 [0, m+n]

// 贪心 + 单调栈
// 1. 从 nums1 中选出 i 个数，从 nums2 中选出 k - i 个数，使得这两个数组拼接后的数组最大
// 2. 遍历 i 的取值，找到使拼接后的数组最大的 i 值
// 3. 使用单调栈来维护 nums1 和 nums2 的子序列，使得子序列中的元素从大到小排列
// 4. 将两个子序列合并，得到最终的数组

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	m, n := len(nums1), len(nums2)
	if k > m+n {
		return []int{}
	}
	ans := make([]int, 0, k)
	for i := 0; i <= k; i++ {
		if i > m {
			continue
		}
		if k-i > n {
			continue
		}
		sub1 := maxSubsequence(nums1, i)
		sub2 := maxSubsequence(nums2, k-i)
		merged := merge(sub1, sub2)
		if len(ans) == 0 || greater(merged, 0, ans, 0) {
			ans = merged
		}
	}
	return ans
}

func maxSubsequence(nums []int, k int) []int {
	stack := []int{}
	for i, num := range nums {
		for len(stack) > 0 && stack[len(stack)-1] < num && len(stack)+len(nums)-i > k {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	return stack[:k]
}

func merge(sub1, sub2 []int) []int {
	res := make([]int, 0, len(sub1)+len(sub2))
	i, j := 0, 0
	for i < len(sub1) && j < len(sub2) {
		if greater(sub1, i, sub2, j) {
			res = append(res, sub1[i])
			i++
		} else {
			res = append(res, sub2[j])
			j++
		}
	}
	res = append(res, sub1[i:]...)
	res = append(res, sub2[j:]...)
	return res
}

func greater(nums1 []int, i int, nums2 []int, j int) bool {
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			return true
		}
		if nums1[i] < nums2[j] {
			return false
		}
		i++
		j++
	}
	return len(nums1)-i > len(nums2)-j
}
