package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var slice = make([]int, 30)
	for i := 0; i < 30; i++ {
		slice[i] = rand.Intn(999)
	}
	fmt.Println(slice)
	sort.Slice(slice, func(a, b int) bool { return slice[a] < slice[b] })
	fmt.Println(slice)

}
