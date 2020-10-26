package factory

import "fmt"

type BlackHuman2 struct {
}

func (human *BlackHuman2) GetColor() {
	fmt.Println("balck")
}

func init() {
	Register("black", func() BaseFactory {
		return new(BlackHuman2)
	})
}
