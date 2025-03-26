# AlgoSphere

AlgoSphere 是一个专注于算法学习和实践的 Java 项目。本项目使用 JDK 17 和 Gradle 构建，包含了常见的数据结构和算法的实现。

## 项目结构

项目按照算法类型组织代码，每个算法类型都有其独立的包和说明文档：

### 基础数据结构
- [数组算法](src/main/java/com/anner/algosphere/array/README.md)
  - 数组旋转、反转、去重、排序等
- [链表算法](src/main/java/com/anner/algosphere/linkedlist/README.md)
  - 链表反转、环检测、合并、排序等
- [树算法](src/main/java/com/anner/algosphere/tree/README.md)
  - 树的遍历、构建、平衡、路径等
- [堆算法](src/main/java/com/anner/algosphere/heap/README.md)
  - 最大堆、最小堆、堆排序等
- [哈希算法](src/main/java/com/anner/algosphere/hash/README.md)
  - 哈希表、冲突解决、布隆过滤器等

### 搜索算法
- [二分查找](src/main/java/com/anner/algosphere/binary_search/README.md)
  - 有序数组查找、旋转数组查找等
- [深度优先搜索](src/main/java/com/anner/algosphere/dfs/README.md)
  - 迷宫寻路、拓扑排序等
- [广度优先搜索](src/main/java/com/anner/algosphere/bfs/README.md)
  - 最短路径、岛屿数量等

### 排序算法
- [排序算法](src/main/java/com/anner/algosphere/sorting/README.md)
  - 冒泡排序、选择排序、插入排序
  - 快速排序、归并排序、堆排序
  - 计数排序、基数排序

### 其他算法
- [递归算法](src/main/java/com/anner/algosphere/recursion/README.md)
  - 阶乘、斐波那契、汉诺塔等
- [动态规划](src/main/java/com/anner/algosphere/dp/README.md)
  - 最长递增子序列、背包问题等
- [图算法](src/main/java/com/anner/algosphere/graph/README.md)
  - 最短路径、最小生成树等
- [字符串算法](src/main/java/com/anner/algosphere/string/README.md)
  - KMP、最长公共子串等
- [数学算法](src/main/java/com/anner/algosphere/math/README.md)
  - 最大公约数、快速幂等

## 学习指南

1. 每个算法目录下都包含：
   - 算法说明文档（README.md）
   - 示例代码
   - 常见应用场景
   - 时间复杂度分析

2. 建议学习顺序：
   - 先掌握基础数据结构（数组、链表、树等）
   - 然后学习基础算法（排序、查找等）
   - 最后学习高级算法（动态规划、图算法等）

3. 实践建议：
   - 先理解算法原理
   - 尝试自己实现
   - 对比不同实现方式
   - 分析时间和空间复杂度

## 环境要求

- JDK 17 或更高版本
- Gradle 8.0 或更高版本

## 贡献指南

欢迎提交 Pull Request 来改进或添加新的算法实现。提交时请确保：
1. 代码符合 Java 编码规范
2. 添加适当的注释和文档
3. 包含必要的测试用例
4. 更新相关的 README 文档 