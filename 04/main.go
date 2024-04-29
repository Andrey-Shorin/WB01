package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
// которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// пример запуска go run main.go -n 10

func worker(id int, channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range channel {
		fmt.Printf("worker №%d data: %d \n", id, i)
		time.Sleep(1 * time.Second)
	}
	return

}
func main() {
	n := flag.Int("n", 3, "amount of workers")
	flag.Parse()
	channel_os := make(chan os.Signal, 1)
	signal.Notify(channel_os, os.Interrupt, syscall.SIGTERM)
	channel_int := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < *n; i++ {
		wg.Add(1)
		go worker(i, channel_int, &wg)
	}
	for {
		select {
		case <-channel_os:
			close(channel_int)
			fmt.Println("\nExit")
			wg.Wait()
			return
		default:
			channel_int <- rand.Int()
		}
	}

}
