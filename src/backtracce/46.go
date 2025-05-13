package backtrace

// 46. 全排列
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

// 回溯法
// 时间复杂度：O(n * n!)
// 空间复杂度：O(n)

func permute(nums []int) [][]int {
	ans := [][]int{}
	backtrace1(nums, []int{}, make([]bool, len(nums)), &ans)
	return ans
}

func backtrace1(nums []int, path []int, used []bool, ans *[][]int) {
	if len(path) == len(nums) {
		*ans = append(*ans, append([]int{}, path...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backtrace1(nums, path, used, ans)
		path = path[:len(path)-1]
		used[i] = false
	}
}

	