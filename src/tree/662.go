package tree

// 662. 二叉树最大宽度
// 给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。


func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 队列中存储节点和其对应的索引
	type pair struct {
		node  *TreeNode
		idx   int
	}
	queue := []pair{{root, 0}}
	maxWidth := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		leftIdx := queue[0].idx
		var rightIdx int
		
		for i := 0; i < levelSize; i++ {
			p := queue[0]
			queue = queue[1:]
			rightIdx = p.idx
			if p.node.Left != nil {
				queue = append(queue, pair{p.node.Left, 2*p.idx})
			}
			if p.node.Right != nil {
				queue = append(queue, pair{p.node.Right, 2*p.idx+1})
			}
		}
		width := rightIdx - leftIdx + 1
		if width > maxWidth {
			maxWidth = width
		}
	}
	return maxWidth
}

// 时间复杂度：O(n)
// 空间复杂度：O(n)