package enum

// 3403 从盒子中找出字典序最大的字符串 I
// 给你一个字符串 word 和一个整数 numFriends。
// 你需要从 word 中选出 numFriends 个字符，使得这些字符的字典序最大。
// 返回选出的字符串。
// 1 <= numFriends <= word.length <= 1000
// word 只包含小写英文字母。

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	n := len(word)
	res := ""
	for i := 0; i < n; i++ {
		// 长度越长, 字典序越大
		res = max(res, word[i:min(i+n-numFriends+1, n)])
	}
	return res
}