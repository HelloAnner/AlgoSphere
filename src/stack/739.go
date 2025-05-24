package stack

// 739. 每日温度
// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指在第 i 天之后，才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 0 来代替。

// 单调栈
// 使用一个栈来存储元素的索引，然后遍历数组，如果当前元素大于栈顶元素，则将栈顶元素弹出，直到栈为空或者栈顶元素大于当前元素，然后将当前元素压入栈。

func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	ans := make([]int, len(temperatures))
	for i, t := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < t {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}

// 时间复杂度：O(n)
// 空间复杂度：O(n)
