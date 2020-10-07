package factory

// 定义工厂接口
type HumansFactory interface {
	CreateHuman() Human
}

// 定义黄种人工厂类并实现CreateHuman方法
type YellowHumanFactory struct {
}

func (h *YellowHumanFactory) CreateHuman() Human {
	return &YellowHuman{}
}

// 定义白种人工厂类并实现CreateHuman方法
type WhiteHumanFactory struct {
}

func (h *WhiteHumanFactory) CreateHuman() Human {
	return &WhiteHuman{}
}

// 定义黑种人工厂类并实现CreateHuman方法
type BlackHumanFactory struct {
}

func (h *BlackHumanFactory) CreateHuman() Human {
	return &BlackHuman{}
}
