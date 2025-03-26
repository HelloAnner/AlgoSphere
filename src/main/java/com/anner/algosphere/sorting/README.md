# 排序算法 (Sorting)

排序算法是将一串数据按照特定顺序进行排列的算法。

## 常见排序算法
1. 冒泡排序 (Bubble Sort)
2. 选择排序 (Selection Sort)
3. 插入排序 (Insertion Sort)
4. 快速排序 (Quick Sort)
5. 归并排序 (Merge Sort)
6. 堆排序 (Heap Sort)
7. 计数排序 (Counting Sort)
8. 基数排序 (Radix Sort)

## 时间复杂度比较
- 冒泡排序：O(n²)
- 选择排序：O(n²)
- 插入排序：O(n²)
- 快速排序：平均 O(nlogn)，最坏 O(n²)
- 归并排序：O(nlogn)
- 堆排序：O(nlogn)
- 计数排序：O(n+k)
- 基数排序：O(d(n+k))

## 示例代码结构
```java
public class Sorting {
    public void quickSort(int[] arr, int low, int high) {
        if (low < high) {
            int pi = partition(arr, low, high);
            quickSort(arr, low, pi - 1);
            quickSort(arr, pi + 1, high);
        }
    }
    
    private int partition(int[] arr, int low, int high) {
        int pivot = arr[high];
        int i = low - 1;
        
        for (int j = low; j < high; j++) {
            if (arr[j] < pivot) {
                i++;
                swap(arr, i, j);
            }
        }
        swap(arr, i + 1, high);
        return i + 1;
    }
}
``` 