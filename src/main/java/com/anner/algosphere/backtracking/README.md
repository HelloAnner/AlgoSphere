# 回溯算法 (Backtracking)

回溯算法是一种通过探索所有可能的候选解来找出所有解的算法。

## 常见应用场景
1. 组合问题
   - 组合总和
   - 子集生成
   - 排列组合
2. 路径问题
   - 迷宫寻路
   - N皇后问题
3. 分割问题
   - 字符串分割
   - IP地址划分

## 示例代码结构
```java
public class BacktrackingAlgorithms {
    public void backtrack(List<Integer> path, List<List<Integer>> result, int[] nums) {
        // 基本情况
        if (isValid(path)) {
            result.add(new ArrayList<>(path));
            return;
        }
        
        // 遍历所有可能的选择
        for (int i = 0; i < nums.length; i++) {
            // 做选择
            path.add(nums[i]);
            
            // 递归
            backtrack(path, result, nums);
            
            // 撤销选择
            path.remove(path.size() - 1);
        }
    }
}
``` 