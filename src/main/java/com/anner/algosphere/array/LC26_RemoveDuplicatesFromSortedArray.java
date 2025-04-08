package com.anner.algosphere.array;

/**
 * 26. 删除有序数组中的重复项
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
 * 
 * 给你一个有序数组 nums，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
 * 
 */
public class LC26_RemoveDuplicatesFromSortedArray {
     public int removeDuplicates(int[] nums) {
        if (nums.length == 0) return 0;
        int k = 0;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != nums[k]) {
                nums[++k] = nums[i];
            }
        }
        return k + 1;
    }
}