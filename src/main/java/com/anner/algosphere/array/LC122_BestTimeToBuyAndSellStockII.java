package com.anner.algosphere.array;

/**
 * 122. 买卖股票的最佳时机 II
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
 * 
 * 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
 * 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。
 * 你也可以先购买，然后在 同一天 出售。
 * 返回 你能获得的 最大 利润 。
 * 
 * 使用贪心算法，只要后一天比前一天价格高，就买入卖出
 * 
 * 对比121题，121题只能持有一股股票，所以只能选择最低点买入，最高点卖出 
 * 122题可以持有多个股票，所以只要后一天比前一天价格高，就买入卖出
 */
public class LC122_BestTimeToBuyAndSellStockII {
    public int maxProfit(int[] prices) {
        int maxProfit = 0;
        for (int i = 1; i < prices.length; i++) {
            if (prices[i] > prices[i - 1]) {
                maxProfit += prices[i] - prices[i - 1];
            }
        }
        return maxProfit;
    }
} 