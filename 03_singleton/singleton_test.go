package singleton

import (
	"sync"
	"testing"
)

const count = 100

func TestSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(count)
	instances := [count]*Singleton{}
	for i := 0; i < count; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < count; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
