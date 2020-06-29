package service

import "fmt"

type car struct {
}

var Instance *car

func init() {
	if Instance == nil {
		Instance = &car{}
	}
}

func (inst *car) Driver() {
	fmt.Println("This is a one instance")
}
