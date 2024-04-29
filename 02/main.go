package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов
// чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
func main() {
	var wg sync.WaitGroup
	numbers := []int{2, 4, 6, 8, 10}
	for _, i := range numbers {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			fmt.Println(number * number)
		}(i)
	}
	wg.Wait()
}
