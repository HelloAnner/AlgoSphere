package dp

// 139. 单词拆分
// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

// dp[i] 代表s的前i个字符是否可以被拆分成字典中的单词
// 状态转移方程：dp[i] = dp[i] || dp[i-len(word)]
// 初始化：dp[0] = true
// 返回：dp[len(s)]

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 || len(wordDict) == 0 {
		return false
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			if i >= len(word) && s[i-len(word):i] == word {
				dp[i] = dp[i] || dp[i-len(word)]
			}
		}
	}
	return dp[len(s)]
}

// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
