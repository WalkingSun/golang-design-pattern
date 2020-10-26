package factory

type BaseFactory interface {
	GetColor()
}

var (
	// 保存注册好的工厂信息
	factoryByName = make(map[string]func() BaseFactory) // func() Class 调用此函数，创建一个类实例，实现的工厂内部结构体会实现Class接口
)

// 注册一个类生成工厂
func Register(name string, factory func() BaseFactory) {
	factoryByName[name] = factory
}

// 根据名称创建对应的类
func Create(name string) BaseFactory {
	if f, ok := factoryByName[name]; ok {
		return f()
	} else {
		panic("name not found")
	}
}
