package factory

import "testing"

func TestSimpleFactory(t *testing.T) {
	factory := new(HumanFactory)
	h1 := factory.CreateHuman("black")
	h1.GetColor()
	h2 := factory.CreateHuman("yellow")
	h2.GetColor()
}

func TestFactoryMethod(t *testing.T) {
	var yellowHuman Human = (&YellowHumanFactory{}).CreateHuman()
	var whitewHuman Human = (&WhiteHumanFactory{}).CreateHuman()
	var blackwHuman Human = (&BlackHumanFactory{}).CreateHuman()
	yellowHuman.GetColor()
	whitewHuman.GetColor()
	blackwHuman.GetColor()
}

func TestAutoloadFactoryMethod(t *testing.T) {
	var yellowHuman Human = Create("yellow")
	var blackwHuman Human = Create("black")
	yellowHuman.GetColor()
	blackwHuman.GetColor()
}
