/**
 * 55. 跳跃游戏
 * https://leetcode.cn/problems/jump-game/
 * 
 * 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 * 判断你是否能够到达最后一个下标。
 * 
 * 使用贪心算法，从后往前遍历，找到能到达最后一个下标的点
 * 
 * 1 2 3 4 5 6 7 8 9 10
 * 2 3 1 1 4 0 0 0 0 0
 * 
 */
public class LC55_CanJump {
    public boolean canJump(int[] nums) {
        int n = nums.length;
        int lastPos = n - 1;
        for (int i = n - 1; i >= 0; i--) {
            if (i + nums[i] >= lastPos) {
                lastPos = i;
            }
        }
        return lastPos == 0;
    }

    /**
     * 从前往后遍历，找到能到达最后一个下标的点
     * 
     * 对比canJump方法，时间复杂度为O(n)
     */
    public boolean canJump2(int[] nums) {
        int n = nums.length;
        int maxPos = 0;
        for (int i = 0; i < n; i++) {
            if (i > maxPos) return false;
            maxPos = Math.max(maxPos, i + nums[i]);
            if (maxPos >= n - 1) return true;
        }
        return false;
    }
}