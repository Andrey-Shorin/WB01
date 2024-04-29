package main

import (
	"fmt"
	"sync"
	"time"
)

type myMap struct {
	data map[int]string
	sync.RWMutex
}

func (s *myMap) Set(key int, value string) {
	s.Lock()
	defer s.Unlock()
	s.data[key] = value
}
func (s *myMap) Get(key int) string {
	s.RLock()
	defer s.RUnlock()
	return s.data[key]
}
func main() {
	var wg sync.WaitGroup

	myMap := myMap{data: make(map[int]string)}
	strings := []string{"one", "neo", "eon", "eno", "noe"}
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		for i, v := range strings {
			myMap.Set(i, v)
			time.Sleep(500 * time.Millisecond)
		}

		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		for i := range strings {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(myMap.Get(i))
		}
		wg.Done()
	}(&wg)

	wg.Wait()
}
