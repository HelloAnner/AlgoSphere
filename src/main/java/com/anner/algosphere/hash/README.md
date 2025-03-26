# 哈希算法 (Hash)

哈希算法是将任意长度的数据映射到固定长度的数据的算法。

## 常见哈希算法
1. 哈希表实现
2. 哈希冲突解决
   - 开放地址法
   - 链地址法
3. 一致性哈希
4. 布隆过滤器
5. 哈希函数设计
6. 哈希查找
7. 哈希排序
8. 哈希统计

## 示例代码结构
```java
public class HashAlgorithms {
    // 哈希表实现
    public class HashTable<K, V> {
        private static final int DEFAULT_CAPACITY = 16;
        private Entry<K, V>[] table;
        private int size;
        
        public HashTable() {
            table = new Entry[DEFAULT_CAPACITY];
            size = 0;
        }
        
        // 插入键值对
        public void put(K key, V value) {
            int index = hash(key);
            if (table[index] == null) {
                table[index] = new Entry<>(key, value);
                size++;
            } else {
                // 处理冲突
                Entry<K, V> entry = table[index];
                while (entry.next != null) {
                    if (entry.key.equals(key)) {
                        entry.value = value;
                        return;
                    }
                    entry = entry.next;
                }
                entry.next = new Entry<>(key, value);
                size++;
            }
        }
        
        // 获取值
        public V get(K key) {
            int index = hash(key);
            Entry<K, V> entry = table[index];
            while (entry != null) {
                if (entry.key.equals(key)) {
                    return entry.value;
                }
                entry = entry.next;
            }
            return null;
        }
        
        // 哈希函数
        private int hash(K key) {
            return Math.abs(key.hashCode() % table.length);
        }
        
        // 哈希表条目
        private static class Entry<K, V> {
            K key;
            V value;
            Entry<K, V> next;
            
            Entry(K key, V value) {
                this.key = key;
                this.value = value;
                this.next = null;
            }
        }
    }
}
``` 