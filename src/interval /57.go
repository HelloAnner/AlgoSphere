package interval

// 57. 插入区间
// 给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
// 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。


func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	ans := [][]int{}
	
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] < newInterval[0] {
			ans = append(ans, intervals[i])
		} else if intervals[i][0] > newInterval[1] { // 当前区间在新区间右侧
			ans = append(ans, newInterval)
			newInterval = intervals[i]
		} else { // 当前区间与新区间有重叠
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[i][1])
		}
	}
	ans = append(ans, newInterval)
	return ans
}

// 时间复杂度：O(n)
// 空间复杂度：O(1)
