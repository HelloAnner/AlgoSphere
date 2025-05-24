package bit

// 201. 数字范围按位与
// 给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字的 按位与 的结果（包含 left 和 right 本身）。

func rangeBitwiseAnd(left int, right int) int {
	// 找到公共前缀
	shift := 0
	// 当left和right不相等时，不断右移，直到找到公共前缀
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	// 将公共前缀左移回原来的位置
	return left << shift
}

// 时间复杂度：O(log n)，其中 n 是范围内的数字个数
// 空间复杂度：O(1)

/* 解题思路：
1. [left, right] 范围内所有数字按位与的结果，其实就是所有数字的二进制表示的公共前缀
2. 例如：[5, 7]
   5 = 101
   6 = 110
   7 = 111
   结果 = 100 (4)
3. 通过不断右移，直到left和right相等，就找到了公共前缀
4. 最后将公共前缀左移回原来的位置即可
*/