package tree

// 106. 从中序与后序遍历序列构造二叉树
// 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历，postorder 是二叉树的后序遍历，请你构造并返回这棵二叉树。


func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}

	index := 0
	for i, v := range inorder {
		if v == root.Val {
			index = i
			break
		}
	}

	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])

	return root
}

// 时间复杂度：O(n)
// 空间复杂度：O(n)


