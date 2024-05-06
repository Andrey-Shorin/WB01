package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func BinarySearch(arr []int, val int) int {
	left, right := 0, len(arr)
	mid := (left + right) / 2

	for left < right {
		if arr[mid] == val {
			return mid
		} else if arr[mid] < val {
			left = mid
		} else {
			right = mid
		}
		mid = (left + right) / 2
	}
	return -1
}

func main() {
	var slice = make(sort.IntSlice, 30, 30)
	for i := 0; i < 30; i++ {
		slice[i] = rand.Intn(999)
	}
	x := slice[rand.Intn(30)]
	sort.Slice(slice, func(a, b int) bool { return slice[a] < slice[b] })
	fmt.Println(slice)
	find := slice.Search(x)
	y := BinarySearch(slice, x)
	fmt.Println("find ", slice[find], " on ", find, "search for ", x)
	fmt.Println("find ", slice[y], " on ", y, "search for ", x)

}
