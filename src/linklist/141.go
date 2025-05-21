package linklist

// 141. 环形链表
// 给定一个链表，判断链表中是否有环。

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

// 时间复杂度：O(n)
// 空间复杂度：O(1)
