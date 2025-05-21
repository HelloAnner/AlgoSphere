package linklist

// 24. 两两交换链表中的节点
// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}
	pre, cur := dummy, head
	
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = cur
		pre.Next = next
		pre = cur
		cur = cur.Next
	}

	return dummy.Next
}

// 时间复杂度：O(n)
// 空间复杂度：O(1)
