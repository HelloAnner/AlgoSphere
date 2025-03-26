# 递归算法 (Recursion)

递归是一种通过函数调用自身来解决问题的方法。递归算法通常包含两个部分：
1. 基本情况（Base Case）：递归终止的条件
2. 递归情况（Recursive Case）：将问题分解为更小的子问题

## 常见应用场景
- 阶乘计算
- 斐波那契数列
- 汉诺塔问题
- 二叉树遍历
- 目录遍历

## 示例代码结构
```java
public class RecursionExample {
    // 基本情况
    if (baseCase) {
        return baseValue;
    }
    
    // 递归情况
    return recursiveCall(smallerProblem);
}
``` 