package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var slice = make(sort.IntSlice, 30, 30)
	for i := 0; i < 30; i++ {
		slice[i] = rand.Intn(999)
	}
	x := slice[rand.Intn(30)]
	sort.Slice(slice, func(a, b int) bool { return slice[a] < slice[b] })
	fmt.Println(slice)
	find := slice.Search(x)
	fmt.Println("find ", slice[find], " on ", find, "search for ", x)

}
