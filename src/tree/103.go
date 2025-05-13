package tree

// 103. 二叉树的锯齿形层序遍历
// 给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
// 广度优先遍历


func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	result := [][]int{}
	leftToRight := true
	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 添加到结果中
		if leftToRight {
			result = append(result, level)
		} else {
			result = append(result, reverse(level))
		}
		leftToRight = !leftToRight
	}
	return result
}

func reverse(level []int) []int {
	for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
		level[i], level[j] = level[j], level[i]
	}
	return level
}
