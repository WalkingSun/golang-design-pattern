# 设计模式
使用Go语言实现设计模式的实例代码。

## 单一职责原则
就一个类而言,应该仅有一个引起 它变化的原因。

如果一个类承担的职责过多，就等于把职责耦合在一起，一个职责的变化可能会削弱或者抑制这个类完成其他职责的能力。这种耦合会导致脆弱的设计，当变化发生时，设计会遭受到意想不到的破环。

软件设计真正要做的许多内容，就是发现职责并把这些职责相互分离。

如果你能够想到多于一个的动机改变一个类，那么这个类就具有多于一个类的职责。

但在软件开发中对于职责的分离是比较抽象，没有固定的标准，要看实际项目中的使用和对于代码结构上分离的感知。比如是书籍类，可以分为属性、动作两个类，动作又可以切分curd类，那岂不是类文件越多的情况。

## 里氏替换原则
面向对象继承优点：
- 代码共享，减少创建类的工作量，每个子类都拥有父类的方法和属性；
- 提供代码的重用性；

缺点：
- 继承是侵入性的。只要继承，就必须拥有父类的所有属性和方法；降低代码的灵活性。子类必须拥有父类的属性和方法，让子类自由的世界中多了些约束；
- 增强了耦合性。当父类的常量、变量和方法被修改时，需要考虑子类的修改，而且在缺乏规范的环境下，这种修改可能带来非常糟糕的结果——大段的代码需要重构。

里氏替换原则：
Functions that use pointers or references to base classes mustbe able to use objects of derived classes without knowing it.（所有引用基类的地方必须能透明地使用其子类的对象。）

通俗点讲，只要父类能出现的地方子类就可以出现，而且替换为子类也不会产生任何错误或异常，使用者可能根本就不需要知道是父类还是子类



## 单例模式
[单例模式-懒汉模式](https://github.com/WalkingSun/DesignPattern/tree/master/singleton/singleton02.go)

[单例模式-饿汉模式](https://github.com/WalkingSun/DesignPattern/tree/master/singleton/singleton.go)

[单例模式-锁机制](https://github.com/WalkingSun/DesignPattern/tree/master/singleton/singleton03.go)

## 工厂模式
[工厂模式](https://github.com/WalkingSun/DesignPattern/tree/master/factory)

[简单工厂模式](https://github.com/WalkingSun/DesignPattern/tree/master/factory/simple_factory.go)

[多工厂方法模式](https://github.com/WalkingSun/DesignPattern/tree/master/factory/more_factory.go)
