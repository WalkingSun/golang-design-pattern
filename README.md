# 设计模式
使用Go语言实现设计模式的实例代码。

## 单一职责原则
就一个类而言,应该仅有一个引起 它变化的原因。

如果一个类承担的职责过多，就等于把职责耦合在一起，一个职责的变化可能会削弱或者抑制这个类完成其他职责的能力。这种耦合会导致脆弱的设计，当变化发生时，设计会遭受到意想不到的破环。

软件设计真正要做的许多内容，就是发现职责并把这些职责相互分离。

如果你能够想到多于一个的动机改变一个类，那么这个类就具有多于一个类的职责。

但在软件开发中对于职责的分离是比较抽象，没有固定的标准，要看实际项目中的使用和对于代码结构上分离的感知。比如是书籍类，可以分为属性、动作两个类，动作又可以切分curd类，那岂不是类文件越多的情况。

## 里氏替换原则(LSP)
[设计原则-里氏替换原则](https://github.com/WalkingSun/golang-design-pattern/tree/master/principle#%E9%87%8C%E6%B0%8F%E6%9B%BF%E6%8D%A2%E5%8E%9F%E5%88%99lsp)

## 开闭原则
[设计原则-开放封闭原则](https://github.com/WalkingSun/golang-design-pattern/tree/master/principle#%E5%BC%80%E6%94%BE%E5%B0%81%E9%97%AD%E5%8E%9F%E5%88%99)

## 单例模式
[单例模式-懒汉模式](https://github.com/WalkingSun/DesignPattern/tree/master/singleton/singleton02.go)

[单例模式-饿汉模式](https://github.com/WalkingSun/DesignPattern/tree/master/singleton/singleton.go)

[单例模式-锁机制](https://github.com/WalkingSun/DesignPattern/tree/master/singleton/singleton03.go)

## 工厂模式
[工厂模式](https://github.com/WalkingSun/DesignPattern/tree/master/factory)

[简单工厂模式](https://github.com/WalkingSun/DesignPattern/tree/master/factory/simple_factory.go)

[多工厂方法模式](https://github.com/WalkingSun/DesignPattern/tree/master/factory/more_factory.go)
