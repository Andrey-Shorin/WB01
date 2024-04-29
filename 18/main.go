package main

import (
	"fmt"
	"sync"
	"time"
)

type Container struct {
	counter int
	sync.Mutex
}

func (container *Container) up() {
	container.Lock()
	defer container.Unlock()
	container.counter += 1

}

func (container *Container) Getcounter() int {
	return container.counter
}
func main() {

	var wt sync.WaitGroup
	container := Container{}
	for i := 0; i < 100; i++ {

		wt.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				container.up()
				time.Sleep(500 * time.Millisecond)
			}
			wt.Done()
		}()
	}
	wt.Wait()
	fmt.Println(container.Getcounter())

}
