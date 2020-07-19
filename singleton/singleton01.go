package singleton

import "fmt"

type Singleton struct {
	S string
}

var Instance *Singleton

func init() {
	if Instance == nil {
		Instance = &Singleton{}
	}
}

func (inst *Singleton) Show() {
	fmt.Println("This is a one instance")
}
