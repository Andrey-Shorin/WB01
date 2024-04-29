package main

import (
	"fmt"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
func main() {

	var N time.Duration = 5
	ch := make(chan int)
	timeout := time.After(N * time.Second)
	go in(ch)
	go out(ch)
	<-timeout
	fmt.Println("End")

}

func in(ch <-chan int) {
	for val := range ch {
		fmt.Println("Resive ", val)
	}
}

func out(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
}
