# 广度优先搜索 (BFS)

广度优先搜索是一种用于遍历或搜索树或图的算法。它从根节点开始，逐层遍历所有节点。

## 基本思想
1. 从根节点开始
2. 访问所有相邻节点
3. 将访问过的节点标记为已访问
4. 对每个相邻节点重复以上步骤

## 常见应用场景
- 最短路径问题
- 社交网络中的好友关系
- 网络爬虫
- 迷宫最短路径
- 岛屿数量问题

## 示例代码结构
```java
public class BFS {
    public void bfs(TreeNode root) {
        if (root == null) return;
        
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        
        while (!queue.isEmpty()) {
            TreeNode node = queue.poll();
            
            // 处理当前节点
            process(node);
            
            // 将子节点加入队列
            if (node.left != null) queue.offer(node.left);
            if (node.right != null) queue.offer(node.right);
        }
    }
}
``` 