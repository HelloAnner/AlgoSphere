package math

// 2929 给小朋友们分糖果 II
// 给n个小朋友分糖果，每个小朋友最多能分到limit个糖果，
// 每个小朋友至少分到一个糖果，问有多少种分法。


func distributeCandies(n int, limit int) int64 {
    var count int64 = 0
    // 遍历所有可能的分配方案
    for a := 0; a <= min(n, limit); a++ {
        remaining := n - a
        if remaining < 0 {
            continue
        }
        // 计算b和c的分配方案数
        minBC := max(0, remaining - limit)
        maxBC := min(limit, remaining)
        if minBC > maxBC {
            continue
        }
        count += int64(maxBC - minBC + 1)
    }
    return count
}