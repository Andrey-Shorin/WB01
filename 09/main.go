package main

import "fmt"

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x)
//из массива, во второй — результат операции x*2, после чего данные из второго
//канала должны выводиться в stdout.

func main() {
	ch := Read(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	pow2nums := PowFromChan(ch)
	for v := range pow2nums {
		fmt.Print(v, " ")
	}
}
func Read(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func PowFromChan(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}
