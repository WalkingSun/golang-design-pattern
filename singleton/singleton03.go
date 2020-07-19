package singleton

import "sync"

type Singleton03 struct {
	S string
}

var singleton03 *Singleton03
var once sync.Once

func GetInstances() *Singleton03 {
	once.Do(func() {
		singleton03 = &Singleton03{}
	})
	return singleton03
}
