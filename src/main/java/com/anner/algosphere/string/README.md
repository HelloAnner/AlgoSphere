# 字符串算法 (String)

字符串算法是处理字符串相关问题的算法集合。

## 常见字符串算法
1. 字符串匹配
   - 暴力匹配
   - KMP 算法
   - Boyer-Moore 算法
2. 最长公共子串
3. 最长回文子串
4. 字符串编辑距离
5. 字符串压缩
6. 字符串反转
7. 字符串排列组合

## 示例代码结构
```java
public class StringAlgorithms {
    // KMP 算法示例
    public int KMP(String text, String pattern) {
        int[] lps = computeLPSArray(pattern);
        int i = 0; // text 的索引
        int j = 0; // pattern 的索引
        
        while (i < text.length()) {
            if (pattern.charAt(j) == text.charAt(i)) {
                i++;
                j++;
            }
            
            if (j == pattern.length()) {
                return i - j; // 找到匹配
            } else if (i < text.length() && pattern.charAt(j) != text.charAt(i)) {
                if (j != 0) {
                    j = lps[j - 1];
                } else {
                    i++;
                }
            }
        }
        
        return -1; // 未找到匹配
    }
    
    private int[] computeLPSArray(String pattern) {
        int[] lps = new int[pattern.length()];
        int len = 0;
        int i = 1;
        
        while (i < pattern.length()) {
            if (pattern.charAt(i) == pattern.charAt(len)) {
                len++;
                lps[i] = len;
                i++;
            } else {
                if (len != 0) {
                    len = lps[len - 1];
                } else {
                    lps[i] = 0;
                    i++;
                }
            }
        }
        
        return lps;
    }
}
``` 