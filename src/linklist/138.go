package linklist

// 138. 复制带随机指针的链表
// 给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。

// 思路:
// 1. 遍历链表，复制每个节点，并将其插入到原节点的后面
// 2. 遍历链表，复制随机指针
// 3. 遍历链表，拆分链表

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = &Node{Val: cur.Val}
		cur.Next.Next = next
		cur = next
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	cur = head	
	newHead := head.Next
	for cur != nil {
		temp := cur.Next
		cur.Next = temp.Next
		if temp.Next != nil {
			temp.Next = temp.Next.Next
		}
		cur = cur.Next
	}

	return newHead
}

// 时间复杂度：O(n)
// 空间复杂度：O(1)
