# 位运算算法 (Bit Manipulation)

位运算算法是利用计算机中数字的二进制表示进行运算的算法。

## 常见应用场景
1. 位运算技巧
   - 判断奇偶
   - 交换两数
   - 判断2的幂
2. 位运算应用
   - 汉明距离
   - 只出现一次的数字
   - 二进制中1的个数

## 示例代码结构
```java
public class BitManipulationAlgorithms {
    // 计算汉明距离
    public int hammingDistance(int x, int y) {
        int xor = x ^ y;
        int distance = 0;
        
        while (xor != 0) {
            distance += xor & 1;
            xor >>= 1;
        }
        
        return distance;
    }
    
    // 判断是否是2的幂
    public boolean isPowerOfTwo(int n) {
        return n > 0 && (n & (n - 1)) == 0;
    }
}
``` 