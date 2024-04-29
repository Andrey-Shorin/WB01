package main

import (
	"fmt"
	"time"
)

func sleep(sec time.Duration) {
	<-time.After(sec)
}

func main() {
	fmt.Println("sleep for 5 seconds")
	sleep(5 * time.Second)
	fmt.Println("main() awake")
}
