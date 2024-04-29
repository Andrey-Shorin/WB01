package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	//1 закрытие канала
	ch := make(chan int)
	go closeChan(ch)
	time.Sleep(500 * time.Millisecond)
	close(ch)

	//2 сигнальный канал
	ch1 := make(chan int)
	ch2 := make(chan int)
	go chan2(ch1, ch2)
	time.Sleep(500 * time.Millisecond)
	ch2 <- 1
	time.Sleep(500 * time.Millisecond)

	//3 context
	ctx, cancel := context.WithCancel(context.Background())
	go cancelChan(ctx)
	time.Sleep(500 * time.Millisecond)
	cancel()
	time.Sleep(500 * time.Millisecond)

	//4 write
	ch2 = make(chan int)
	go chanWrite(ch2)
	time.Sleep(500 * time.Millisecond)
	<-ch2
	time.Sleep(500 * time.Millisecond)

	//5 wait
	var wg sync.WaitGroup
	wg.Add(1)
	go wait(&wg)
	wg.Done()
	fmt.Println("all ended")
}

func closeChan(in <-chan int) {
	for range in {
	}
	fmt.Println("chanal 1 closed")
}
func chan2(in <-chan int, signal <-chan int) {
	_ = in
	<-signal
	fmt.Println("chanal 2 closed")
}

func cancelChan(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("chanal 3 closed")
}

func chanWrite(in chan<- int) {
	in <- 1
	fmt.Println("chanal 4 closed")
}
func wait(wg *sync.WaitGroup) {
	wg.Wait()
	fmt.Println("chanal 5 closed")
}
