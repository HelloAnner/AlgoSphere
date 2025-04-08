package com.anner.algosphere.array;

import java.util.ArrayList;
import java.util.List;

/**
 * 228. 汇总区间
 * https://leetcode.cn/problems/summary-ranges/
 * 
 * 给定一个 无重复元素 的 有序 整数数组 nums 。
 * 返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。
 */
public class LC228_SummaryRanges {
    public List<String> summaryRanges(int[] nums) {
        List<String> result = new ArrayList<>();
        if (nums == null || nums.length == 0) {
            return result;
        }

        int n = nums.length;
        int i = 0;
        while (i < n) {
            int start = nums[i];
            while (i + 1 < n && nums[i + 1] == nums[i] + 1) {
                i++;
            }
            int end = nums[i];
            if (start == end) {
                result.add(String.valueOf(start));
            } else {
                result.add(start + "->" + end);
            }
            i++;
        }
        return result;
    }
} 