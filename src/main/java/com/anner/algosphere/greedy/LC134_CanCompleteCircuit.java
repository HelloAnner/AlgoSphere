/**
 * 134. 加油站
 * 
 * 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。   
 * 
 * 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
 * 
 * 给定两个整数数组 gas 和 cost ，如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
 */

 public class LC134_CanCompleteCircuit {
    public int canCompleteCircuit(int[] gas, int[] cost) {
        int n = gas.length;
        int totalGas = 0, totalCost = 0;
        // 计算总油量和总消耗量
        for (int i = 0; i < n; i++) {
            totalGas += gas[i];
            totalCost += cost[i];
        }
        if (totalGas < totalCost) {
            return -1;
        }
        // 从0号加油站开始，计算当前油量
        int currentGas = 0, start = 0;
        for (int i = 0; i < n; i++) {
            currentGas += gas[i] - cost[i];
            // 如果当前油量小于0，则从下一个加油站开始
            if (currentGas < 0) {
                start = i + 1;
                currentGas = 0;
            }
        }
        return start;
    }
 }