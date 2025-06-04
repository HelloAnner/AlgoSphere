package heap

import (
	"container/heap"
	"sort"
)

// 295. 数据流中的中位数
// 中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
// 例如，
// [2,3,4] 的中位数是 3
// [2,3] 的中位数是 (2 + 3) / 2 = 2.5
// 设计一个支持以下两种操作的数据结构：
// void addNum(int num) - 从数据流中添加一个整数到数据结构中。
// double findMedian() - 返回目前所有元素的中位数。
type MedianFinder struct {
    left  hp // 入堆的元素取相反数，变成最大堆
    right hp // 最小堆
}

func Constructor() (_ MedianFinder) { return }

func (mf *MedianFinder) AddNum(num int) {
    if mf.left.Len() == mf.right.Len() {
        heap.Push(&mf.right, num)
        heap.Push(&mf.left, -heap.Pop(&mf.right).(int))
    } else {
        heap.Push(&mf.left, -num)
        heap.Push(&mf.right, -heap.Pop(&mf.left).(int))
    }
}

func (mf *MedianFinder) FindMedian() float64 {
    if mf.left.Len() > mf.right.Len() {
        return float64(-mf.left.IntSlice[0])
    }
    return float64(mf.right.IntSlice[0]-mf.left.IntSlice[0]) / 2.0
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
