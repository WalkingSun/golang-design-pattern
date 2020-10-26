package factory

import "fmt"

type YellowHuman2 struct {
}

func (human *YellowHuman2) GetColor() {
	fmt.Println("yellow")
}

func init() {
	Register("yellow", func() BaseFactory {
		return new(YellowHuman2)
	})
}
