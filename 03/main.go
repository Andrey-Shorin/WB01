package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
//квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	numbers := []int{2, 4, 6, 8, 10}
	sumnumbers := 0
	for _, i := range numbers {

		wg.Add(1)
		go func(number int) {

			defer wg.Done()

			mu.Lock()

			defer mu.Unlock()
			sumnumbers += number * number
		}(i)

	}

	wg.Wait()
	fmt.Println(sumnumbers)
}
