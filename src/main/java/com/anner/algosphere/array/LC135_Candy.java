package com.anner.algosphere.array;

import java.util.Arrays;

/**
 * 135. 分发糖果
 * 
 * n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
 * 
 * 你需要按照以下要求，给这些孩子分发糖果：
 * 
 * 每个孩子至少分配到 1 个糖果。
 * 相邻两个孩子评分更高的孩子会获得更多的糖果。
 * 
 * 请你给每个孩子分发糖果，计算并返回需要准备的最少糖果数目。
 */
public class LC135_Candy {
    public int candy(int[] ratings) {
        int n = ratings.length;
        int[] candies = new int[n];
        // 初始化每个孩子至少分配到1个糖果
        Arrays.fill(candies, 1);
        // 从左到右遍历，确保右边评分更高的孩子获得更多糖果
        for (int i = 1; i < n; i++) {
            if (ratings[i] > ratings[i - 1]) {
                candies[i] = candies[i - 1] + 1;
            }
        }
        // 从右到左遍历，确保左边评分更高的孩子获得更多糖果
        for (int i = n - 2; i >= 0; i--) {
            if (ratings[i] > ratings[i + 1]) {
                candies[i] = Math.max(candies[i], candies[i + 1] + 1);
            }
        }
        // 计算总糖果数
        int totalCandies = 0;
        for (int candy : candies) {
            totalCandies += candy;
        }
        return totalCandies;
    }
}
