# 设计模式 (Design)

设计模式是软件开发中常见问题的解决方案。

## 常见设计模式
1. 创建型模式
   - 单例模式
   - 工厂模式
   - 建造者模式
2. 结构型模式
   - 适配器模式
   - 装饰器模式
   - 代理模式
3. 行为型模式
   - 观察者模式
   - 策略模式
   - 命令模式

## 示例代码结构
```java
// 单例模式示例
public class Singleton {
    private static volatile Singleton instance;
    
    private Singleton() {}
    
    public static Singleton getInstance() {
        if (instance == null) {
            synchronized (Singleton.class) {
                if (instance == null) {
                    instance = new Singleton();
                }
            }
        }
        return instance;
    }
}

// 工厂模式示例
public interface Product {
    void operation();
}

public class ConcreteProduct implements Product {
    @Override
    public void operation() {
        // 具体实现
    }
}

public class Factory {
    public Product createProduct() {
        return new ConcreteProduct();
    }
}
``` 