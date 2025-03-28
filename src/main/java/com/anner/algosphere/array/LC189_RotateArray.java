/**
 * 189. 轮转数组
 * https://leetcode.cn/problems/rotate-array/
 * 
 * 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
 * 
 * 1 2 3 4 5 6 7 , k = 3
 * 整体翻转 7 6 5 4 3 2 1
 * 前k个翻转 5 6 7 4 3 2 1
 * 后n-k个翻转 5 6 7 1 2 3 4
 */
public class LC189_RotateArray {
    public void rotate(int[] nums, int k) {
        int n = nums.length;
        k = k % n;
        // 先整体反转
        reverse(nums, 0, n - 1);
        // 再反转前k个
        reverse(nums, 0, k - 1);
        // 再反转后n-k个
        reverse(nums, k, n - 1);
    }

    private void reverse(int[] nums, int start, int end) {
        while (start < end) {
            int temp = nums[start];
            nums[start] = nums[end];
            nums[end] = temp;
            start++;
            end--;
        }
    }
}