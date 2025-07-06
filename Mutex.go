package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu      sync.Mutex
	counter map[string]int
}

func (c *Container) add(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter[name]++
}

func main() {
	container := &Container{}
	container.counter = map[string]int{"a": 0, "b": 0, "c": 0}

	var wg sync.WaitGroup

	doIncrement := func(name string, num int) {
		defer wg.Done()
		for i := 0; i < num; i++ {
			container.add(name)
		}
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(container.counter)
}
