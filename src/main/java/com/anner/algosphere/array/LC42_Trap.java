package com.anner.algosphere.array;

/**
 * 42. 接雨水
 * 
 * 链接: https://leetcode.cn/problems/trapping-rain-water/
 * 
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 * 
 * 示例 1：
 * 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出：6
 * 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
 *
 */
public class LC42_Trap {
    public int trap(int[] height) {
        int left = 0, right = height.length - 1;
        int leftMax = 0, rightMax = 0;
        int water = 0;
        // 使用双指针法，从两端向中间移动
        while (left < right) {
            // 更新左边的最大高度
            leftMax = Math.max(leftMax, height[left]);
            // 更新右边的最大高度
            rightMax = Math.max(rightMax, height[right]);
            // 如果左边的高度小于右边的高度，则计算左边的雨水量
            if (height[left] < height[right]) {
                water += leftMax - height[left];
                left++;
            } else {
                // 如果左边的高度大于右边的高度，则计算右边的雨水量
                water += rightMax - height[right];
                right--;
            }
        }
        return water;
    }
}
