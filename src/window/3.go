package window

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/?envType=study-plan-v2&envId=top-interview-150
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// 思路: 滑动窗口
// 1. 使用两个指针 left 和 right 表示当前窗口的左右边界
// 2. 使用一个 map 来记录当前窗口内每个字符出现的次数
// 3. 移动 right 指针，扩大窗口，更新 map
// 4. 当 map 中某个字符出现的次数大于 1 时，移动 left 指针，缩小窗口，更新 map
// 5. 记录最大的窗口长度

func LengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	window := make(map[byte]int)
	maxLen := 0

	for right < len(s) {
		window[s[right]]++
		for window[s[right]] > 1 {
			window[s[left]]--
			left++
		}
		maxLen = max(maxLen, right-left+1)
		right++
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
