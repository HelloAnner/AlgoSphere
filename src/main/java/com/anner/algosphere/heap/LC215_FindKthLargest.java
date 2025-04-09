package com.anner.algosphere.heap;

import java.util.PriorityQueue;

/**
 * 215. 数组中的第K个最大元素
 * https://leetcode.cn/problems/kth-largest-element-in-an-array/
 * 
 * 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
 * 
 * 维护一个大小为k的最小堆，堆顶元素即为第k个最大的元素，原因：
 * 1. 堆顶元素是堆中最小的元素，如果堆中元素小于k，则将元素加入堆中
 * 2. 如果堆中元素大于k，则将堆中最小的元素移除，因为堆顶元素是最小的元素，所以堆顶元素即为第k个最大的元素
 * 
 * 相当于不断的把最小元素不断的淘汰，最后堆顶元素即为第k个最大的元素
 * 
 * 时间复杂度：O(nlogk)
 * 空间复杂度：O(k)
 */
public class LC215_FindKthLargest {
    public int findKthLargest(int[] nums, int k) {
        PriorityQueue<Integer> minHeap = new PriorityQueue<>();
        for (int num : nums) {
            minHeap.offer(num);
            if (minHeap.size() > k) {
                minHeap.poll();
            }
        }
        return minHeap.peek();
    }

}
