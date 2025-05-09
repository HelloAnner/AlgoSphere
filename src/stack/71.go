package stack

// 71. Simplify Path 简化路径
// 给你一个字符串 path ，表示指向某一文件或目录的 Unix 风格 绝对路径 （以 '/' 开头），请你将其转化为更加简洁的规范路径。
// 在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..） 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。
// 任意多个连续的斜杠（即，'//'）都被视为单个斜杠 '/' 。 对于此问题，任何其他格式的点（例如，'...'）均被视为文件/目录名称。

import "strings"

func simplifyPath(path string) string {
	stack := []string{}
	for _, dir := range strings.Split(path, "/") {
		if dir == "" || dir == "." {
			continue
		}
		if dir == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, dir)
		}
	}
	return "/" + strings.Join(stack, "/")
}

// 时间复杂度：O(n)
// 空间复杂度：O(n)
