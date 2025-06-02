package random

import "math/rand"

// 384. 打乱数组
// 给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。打乱后，数组的所有排列应该是 等可能 的。

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	ans := make([]int, n)
	copy(ans, this.nums)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

