package interval

import "sort"

// 56. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{}
	
	for i := 0; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]
		
		if len(ans) > 0 && ans[len(ans)-1][1] >= start {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], end)
		} else {
			ans = append(ans, []int{start, end})
		}
	}

	return ans
}

// 时间复杂度：O(nlogn)
// 空间复杂度：O(logn)
