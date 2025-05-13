package backtrace

// 17. 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

var letterMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ans := []string{}
	backtrace2(digits, 0, "", &ans)
	return ans
}

func backtrace2(digits string, index int, path string, ans *[]string) {
	if index == len(digits) {
		*ans = append(*ans, path)
		return
	}

	digit := digits[index]
	letters := letterMap[digit]
	for _, letter := range letters {
		path += letter
		backtrace2(digits, index+1, path, ans)
		path = path[:len(path)-1]
	}
}
