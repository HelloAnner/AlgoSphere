# 链表算法 (LinkedList)

链表算法是处理链表相关问题的算法集合。

## 常见链表算法
1. 链表反转
2. 链表环检测
3. 链表合并
4. 链表排序
5. 链表删除
6. 链表插入
7. 链表查找
8. 链表重排

## 示例代码结构
```java
public class LinkedListAlgorithms {
    // 链表节点定义
    public static class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }
    }
    
    // 链表反转
    public ListNode reverseList(ListNode head) {
        ListNode prev = null;
        ListNode curr = head;
        
        while (curr != null) {
            ListNode nextTemp = curr.next;
            curr.next = prev;
            prev = curr;
            curr = nextTemp;
        }
        
        return prev;
    }
    
    // 检测链表是否有环
    public boolean hasCycle(ListNode head) {
        if (head == null || head.next == null) {
            return false;
        }
        
        ListNode slow = head;
        ListNode fast = head;
        
        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
            
            if (slow == fast) {
                return true;
            }
        }
        
        return false;
    }
} 