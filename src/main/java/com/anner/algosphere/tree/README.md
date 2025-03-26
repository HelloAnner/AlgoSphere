# 树算法 (Tree)

树算法是处理树结构相关问题的算法集合。

## 常见树算法
1. 树的遍历
   - 前序遍历
   - 中序遍历
   - 后序遍历
   - 层序遍历
2. 树的构建
3. 树的平衡
4. 树的路径
5. 树的深度
6. 树的宽度
7. 树的镜像
8. 树的序列化

## 示例代码结构
```java
public class TreeAlgorithms {
    // 树节点定义
    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }
    
    // 前序遍历
    public void preorderTraversal(TreeNode root) {
        if (root == null) return;
        
        // 处理当前节点
        System.out.print(root.val + " ");
        
        // 递归遍历左子树
        preorderTraversal(root.left);
        
        // 递归遍历右子树
        preorderTraversal(root.right);
    }
    
    // 层序遍历
    public void levelOrderTraversal(TreeNode root) {
        if (root == null) return;
        
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        
        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            for (int i = 0; i < levelSize; i++) {
                TreeNode node = queue.poll();
                System.out.print(node.val + " ");
                
                if (node.left != null) queue.offer(node.left);
                if (node.right != null) queue.offer(node.right);
            }
            System.out.println();
        }
    }
} 