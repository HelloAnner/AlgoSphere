package window

// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/?envType=study-plan-v2&envId=top-interview-150
// 给你一个字符串 s 和字符串数组 words 。words 中所有字符串 长度相同 。
// 返回 s 中所有以 words 中所有字符串为子串的连接串的开始索引。

// 思路: 滑动窗口
// 1. 使用两个指针 left 和 right 表示当前窗口的左右边界
// 2. 使用一个 map 来记录当前窗口内每个字符出现的次数
// 3. 移动 right 指针，扩大窗口，更新 map
// 4. 当 map 中每个字符出现的次数都大于等于 t 中每个字符出现的次数时，移动 left 指针，缩小窗口，更新 map
// 5. 记录最小的窗口长度

func FindSubstring(s string, words []string) (ans []int) {
	 ls, m, n := len(s), len(words), len(words[0])
    for i := 0; i < n && i+m*n <= ls; i++ {
        differ := map[string]int{}
        for j := 0; j < m; j++ {
            differ[s[i+j*n:i+(j+1)*n]]++
        }
        for _, word := range words {
            differ[word]--
            if differ[word] == 0 {
                delete(differ, word)
            }
        }
        for start := i; start < ls-m*n+1; start += n {
            if start != i {
                word := s[start+(m-1)*n : start+m*n]
                differ[word]++
                if differ[word] == 0 {
                    delete(differ, word)
                }
                word = s[start-n : start]
                differ[word]--
                if differ[word] == 0 {
                    delete(differ, word)
                }
            }
            if len(differ) == 0 {
                ans = append(ans, start)
            }
        }
	}
	return ans
}
