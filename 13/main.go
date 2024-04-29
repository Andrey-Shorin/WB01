package main

import "fmt"

func main() {
	a, b := 15, 10
	fmt.Println(a, b)
	b, a = a, b
	fmt.Println(a, b)
}
