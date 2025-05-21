package interval

import "fmt"

// 228. 汇总区间
// 给定一个  无重复元素 的 有序 整数数组 nums 。
// 返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。
// 列表中的每个区间范围 [a,b] 应该按如下格式输出：
// "a->b" ，如果 a != b
// "a" ，如果 a == b

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	ans := []string{}
	
	for i := 0; i < len(nums); i++ {
		start := nums[i]
		for i+1 < len(nums) && nums[i+1] == nums[i]+1 {
			i++
		}
		if start == nums[i] {
			ans = append(ans, fmt.Sprintf("%d", start))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", start, nums[i]))
		}
	}
	return ans
}

// 时间复杂度：O(n)
// 空间复杂度：O(1)
