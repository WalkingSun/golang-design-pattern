package singleton

import (
	"sync"
	"testing"
)

const parCount02 = 1000

func TestSingleton02(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	//fmt.Println( fmt.Sprintf("%p",ins1),fmt.Sprintf("%p",ins2))
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func TestParallelSingleton02(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(parCount02)
	instances := make([]*Singleton02, parCount02, parCount02)
	for i := 0; i < parCount02; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < parCount02; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
