package factory

import "fmt"

// Human 是人类活动特性的抽象
type Human interface {
	GetColor()
}

type BlackHuman struct {
}

type WhiteHuman struct {
}

type YellowHuman struct {
}

func (human *BlackHuman) GetColor() {
	fmt.Println("balck")
}

func (human *WhiteHuman) GetColor() {
	fmt.Println("white")
}

func (human *YellowHuman) GetColor() {
	fmt.Println("yellow")
}

type HumanFactory struct {
}

func (factory HumanFactory) CreateHuman(human string) Human {
	switch human {
	case "black":
		return &BlackHuman{}
	case "white":
		return &WhiteHuman{}
	case "yellow":
		return &YellowHuman{}
	default:

	}
	return nil
}
