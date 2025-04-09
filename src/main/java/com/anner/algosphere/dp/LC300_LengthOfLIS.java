package com.anner.algosphere.dp;

import java.util.Arrays;

/**
 * 300. 最长递增子序列
 * https://leetcode.cn/problems/longest-increasing-subsequence/
 * 
 * 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
 * 
 * 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7]
 * 的子序列。
 * 
 * 动态规划:
 * 1. 状态定义：dp[i] 表示以 nums[i] 结尾的最长递增子序列的长度
 * 2. 状态转移：dp[i] = max(dp[i], dp[j] + 1)，其中 j < i 且 nums[j] < nums[i]
 * 3. 初始化：dp[i] = 1
 * 4. 结果：max(dp)
 * 
 * 时间复杂度：O(n^2)
 * 空间复杂度：O(n)
 */
public class LC300_LengthOfLIS {
    public int lengthOfLIS(int[] nums) {
        int[] dp = new int[nums.length];
        Arrays.fill(dp, 1);
        for (int i = 0; i < nums.length; i++) {
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j]) {
                    dp[i] = Math.max(dp[i], dp[j] + 1);
                }
            }
        }
        return Arrays.stream(dp).max().getAsInt();
    }
}