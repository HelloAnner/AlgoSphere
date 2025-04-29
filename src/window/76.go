package window

// https://leetcode.cn/problems/minimum-window-substring/?envType=study-plan-v2&envId=top-interview-150
// 给你一个字符串 s 、一个字符串 t 。返回 s 中包含 t 所有字符的最小子串。如果 s 中不存在这样的子串，则返回空字符串 "" 。
// 思路: 滑动窗口
// 1. 使用两个指针 left 和 right 表示当前窗口的左右边界
// 2. 使用一个 map 来记录当前窗口内每个字符出现的次数
// 3. 移动 right 指针，扩大窗口，更新 map
// 4. 当 map 中每个字符出现的次数都大于等于 t 中每个字符出现的次数时，移动 left 指针，缩小窗口，更新 map
// 5. 记录最小的窗口长度

import "math"

func MinWindow(s string, t string) string {
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	window := make(map[byte]int)
	left, right := 0, 0
	valid := 0
	minLen := math.MaxInt32
	res := ""
	for right < len(s) {
		window[s[right]]++
		if window[s[right]] == need[s[right]] {
			valid++
		}
		for valid == len(need) {
			if right-left+1 < minLen {
				minLen = right - left + 1
				res = s[left:right+1]
			}
			window[s[left]]--
			if window[s[left]] < need[s[left]] {
				valid--
			}
			left++
		}
		right++
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return res
}
