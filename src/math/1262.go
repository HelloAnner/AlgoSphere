package math

import "sort"

// 1262. 可被三整除的最大和
// 给你一个整数数组 nums，请你找出并返回能被三整除的元素最大和。

// 贪心算法
// 1. 计算数组所有元素的和
// 2. 计算数组所有元素的模3的余数
// 3. 如果数组所有元素的和能被3整除，则返回和
// 4. 如果数组所有元素的和不能被3整除，则需要考虑删除一个元素
// 5. 如果数组所有元素的和不能被3整除，则需要考虑删除两个元素


func maxSumDivThree(nums []int) int {
	sum := 0
	mod1 := []int{}
	mod2 := []int{}

	for _, num := range nums {
		sum += num
		if num % 3 == 1 {
			mod1 = append(mod1, num)
		} else if num % 3 == 2 {
			mod2 = append(mod2, num)
		}
	}

	sort.Ints(mod1)
	sort.Ints(mod2)

	if sum % 3 == 0 {
		return sum
	}

	if sum % 3 == 1 {
		if len(mod1) > 0 {
			sum -= mod1[0]
		}
		if len(mod2) > 1 {
			sum -= mod2[0] + mod2[1]
		}
	}

	return sum
}

