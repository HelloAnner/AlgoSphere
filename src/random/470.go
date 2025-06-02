package random

// 470. 用 Rand7() 实现 Rand10()
// 给定方法 rand7 可生成 [1,7] 范围内的均匀随机整数，试写一个方法 rand10 生成 [1,10] 范围内的均匀随机整数。
// 你只能调用 rand7() 且不能调用其他方法。请不要使用系统的 Math.random() 方法。
	

func rand10() int {
    for {
        row := rand7()
        col := rand7()
        idx := (row - 1) * 7 + col
        if idx <= 40 {
            return 1 + (idx - 1) % 10
        }
        row = idx - 40
        col = rand7()
        idx = (row - 1) * 7 + col
        if idx <= 60 {
            return 1 + (idx - 1) % 10
        }
        row = idx - 60
        col = rand7()
        idx = (row - 1) * 7 + col
        if idx <= 20 {
            return 1 + (idx - 1) % 10
        }
    }
}

func rand7() int {
	return 0
}