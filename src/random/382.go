package random

import "math/rand"

// 382 链表随机节点
// 给定一个单链表，随机选择链表的一个节点，并返回相应的节点值。每个节点 被选中的概率一样 。
// 如果链表非常大且长度未知，该怎么处理？
// 你能否在不使用额外空间的情况下解决此问题？

// 蓄水池抽样算法
// 从链表头开始，遍历整个链表，对遍历到的第 i 个节点，随机选择区间 [0,i) 内的一个整数，如果其等于 0，则将答案置为该节点的值，否则答案不变。
// 该算法会保证每个节点的值成为答案的概率均为 1/n，其中 n 为链表的长度。


type ListNode struct {
	Val int
	Next *ListNode
}

type Solution1 struct {
	head *ListNode
}

func Constructor1(head *ListNode) Solution1 {
	return Solution1{head}
}

func (this *Solution1) GetRandom() int {
	ans, i := 0, 0
	for node := this.head; node != nil; node = node.Next {
		i++
		if rand.Intn(i) == 0 {
			ans = node.Val
		}
	}
	return ans
}