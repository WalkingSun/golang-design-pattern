package singleton

type Singleton02 struct {
	S string
}

var singleton *Singleton02

func GetInstance() *Singleton02 {
	if singleton == nil {
		singleton = &Singleton02{}
	}

	return singleton
}
