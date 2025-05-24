package bit

// 136. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

// 异或运算符
// 1. 任何数和0做异或运算，结果仍然是原来的数，即 a ^ 0 = a
// 2. 任何数和其自身做异或运算，结果是0，即 a ^ a = 0
// 3. 异或运算满足交换律和结合律，即 a ^ b ^ a = b ^ a ^ a = b ^ (a ^ a) = b ^ 0 = b

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}