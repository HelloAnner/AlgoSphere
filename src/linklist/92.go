package linklist

// 92. 反转链表 II
// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
// 请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}
	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	
	cur := pre.Next
	for i := left; i < right; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummy.Next
}

// 时间复杂度：O(n)	
// 空间复杂度：O(1)
