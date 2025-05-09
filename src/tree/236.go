package tree

// 236. 二叉树的最近公共祖先
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	return dfs(root, p, q)
}

func dfs(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := dfs(root.Left, p, q)
	right := dfs(root.Right, p, q)
	// 如果左右子树都找到了，则当前节点就是最近公共祖先
	if left != nil && right != nil {
		return root
	}
	// 如果左子树找到了，则返回左子树
	if left != nil {
		return left
	}
	// 如果右子树找到了，则返回右子树
	if right != nil {
		return right
	}
	// 如果左右子树都没有找到，则返回 nil
	return nil
}
