package bfs

// 1298. 你能从盒子里获得的最大糖果数
// 给你 n 个盒子，每个盒子的格式为 [status, candies, keys, containedBoxes] ，其中：
//
// status[i] 整数，如果 box[i] 是开的，那么是 1 ，否则是 0 。
// candies[i] 整数，表示 box[i] 中糖果的数目。
// keys[i] 数组，表示你打开 box[i] 后，可以得到一些盒子的钥匙，每个元素分别为该钥匙对应盒子的下标。
// containedBoxes[i] 整数，表示通过打开 box[i] 后，可以存放的盒子下标，
// 你可以将它们放在任何盒子中。
// 请你按照如下规则，返回可以获得糖果的 最大数目 ：
//
// 在状态为 1 的盒子中，如果其中有糖果，那么你可以获得它。
// 对盒子中的每一个钥匙，如果它对应的盒子是可以打开的，那么你也可以获得那个盒子。
// 你可以随时打开一个盒子，你也可以在拿到盒子里的糖果以后将其放在盒子里。
//
// 示例 1：
// 输入：status = [1,0,1,0], candies = [7,5,4,100], keys = [[],[],[1],[]], containedBoxes = [[1,2],[3],[],[]]
// 输出：16
// 解释：你将获得以下糖果：
// - 从第 0 个盒子开始，你可以获得包含第 1 和 2 个盒子的钥匙。
// - 从第 0 个盒子，你可以获得第 1 个盒子中的糖果。
// - 从第 0 个盒子，你可以获得第 2 个盒子中的糖果。
// - 从第 3 个盒子，你可以获得第 1 个盒子中的糖果。
// 你可以获得糖果的总数目 = 7 + 5 + 4 + 100 = 16 个。

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int)  (ans int) {
    n := len(status)
    visited := make([]bool, n) // 记录盒子是否被访问过
    haveBox := make([]bool, n) // 记录当前拥有的盒子
    haveKey := make([]bool, n) // 记录当前拥有的钥匙
    queue := []int{}
    for _, box := range initialBoxes {
        haveBox[box] = true
    }
    for {
        found := false
        for i := 0; i < n; i++ {
            if haveBox[i] && !visited[i] && (status[i] == 1 || haveKey[i]) {
                visited[i] = true
                found = true
                queue = append(queue, i)
            }
        }
        if !found {
            break
        }
        for len(queue) > 0 {
            box := queue[0]
            queue = queue[1:]
            // 拿糖果
            ans += candies[box]
            // 获得钥匙
            for _, key := range keys[box] {
                haveKey[key] = true
            }
            // 获得盒子
            for _, b := range containedBoxes[box] {
                haveBox[b] = true
            }
        }
    }
    return ans
}