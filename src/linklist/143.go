package linklist

// 143. 重排链表

// 876. 链表的中间结点
func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
    var pre, cur *ListNode = nil, head
    for cur != nil {
        nxt := cur.Next
        cur.Next = pre
        pre = cur
        cur = nxt
    }
    return pre
}

func reorderList(head *ListNode) {
    mid := middleNode(head)
	// 翻转后半部分
    head2 := reverseList(mid)
	// 交替合并两个链表
    for head2.Next != nil {
        nxt := head.Next // A -> B 中的 B
        nxt2 := head2.Next // 1 -> 2 中的 2
        head.Next = head2 // A -> 1
        head2.Next = nxt // 1 -> B
        head = nxt // A -> B
        head2 = nxt2 // 1 -> 2
    }
}