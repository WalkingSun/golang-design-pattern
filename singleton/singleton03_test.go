package singleton

import (
	"sync"
	"testing"
)

func TestSingleton03(t *testing.T) {
	ins1 := GetInstances()
	ins2 := GetInstances()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func TestParallelSingleton03(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*Singleton03{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			instances[index] = GetInstances()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
