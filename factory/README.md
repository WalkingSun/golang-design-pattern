# 工厂模式
Define an interface for creating an object,but let subclasses decide whichclass to instantiate.Factory Method lets a class defer instantiation tosubclasses.
（定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法使一个类的实例化延迟到其子类。）

![image-20200905222329509](https://raw.githubusercontent.com/WalkingSun/WindBlog/gh-pages/images/ws2/image-20200905222329509.png)
抽象产品类Product负责定义产品的共性，实现对事物最抽象的定义；Creator为抽象创建类，也就是抽象工厂，具体如何创建产品类是由具体的实现工厂ConcreteCreator完成的。

## 简单工厂模式
一个模块仅需要一个工厂类，没有必要把它产生出来，使用静态的方法就可以操作；

参考《设计模式之禅》女娲造人列子：
![image-20200905223730728](https://raw.githubusercontent.com/WalkingSun/WindBlog/gh-pages/images/ws2/image-20200905223730728.png)

创建一个对象的时候，调用同一个方法，传入不同的参数就可以返回给我们不同的对象；

工厂模式的优缺点:

优点: 工厂类是整个工厂模式的核心，我们只需要传入给定的信息，就可以创建所需实例，在多人协作的时候，无需知道对象之间的内部依赖，可以直接创建，有利于整个软件体系结构的优化

缺点: 工厂类中包含了所有实例的创建逻辑，一旦这个工厂类出现问题，所有实例都会受到影响，并且，工厂类中生产的产品都基于一个共同的接口，一旦要添加不同种类的产品，这就会增加工厂类的复杂度，将不同种类的产品混合在一起，
违背了单一职责，系统的灵活性和可维护性都会降低，并且当新增产品的时候，必须要修改工厂类，违背了开闭原则『系统对扩展开放，对修改关闭』的原则

