package prefix

// 560. 和为 K 的子数组
// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。


func subarraySum(nums []int, k int) (ans int) {
    s := make([]int, len(nums)+1)
    for i, x := range nums {
        s[i+1] = s[i] + x
    }

    cnt := make(map[int]int, len(s)) // 预分配空间
    for _, sj := range s {
        ans += cnt[sj-k]
        cnt[sj]++
    }
	return ans
}

// 时间复杂度：O(n)
// 空间复杂度：O(n)

