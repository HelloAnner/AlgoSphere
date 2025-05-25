package tree

// LCR 143. 子树判断
// 给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。


func isSubStructure(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if t == nil {
		return false
	}
	if isSameTree(s, t) {
		return true
	}
	if s.Left != nil && isSubStructure(s.Left, t) {	
		return true
	}
	if s.Right != nil && isSubStructure(s.Right, t) {
		return true
	}
	return false
}

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		return true
	}
	if s == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}

// 时间复杂度：O(n)
// 空间复杂度：O(n)
