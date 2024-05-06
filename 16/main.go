package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func quickSort(arr []int, start, fin int) []int {
	if start < fin {
		var p int
		arr, p = partition(arr, start, fin)
		arr = quickSort(arr, start, p-1)
		arr = quickSort(arr, p+1, fin)
	}
	return arr
}

func partition(arr []int, start, fin int) ([]int, int) {
	pivot := arr[fin]
	i := start
	for j := start; j < fin; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[fin] = arr[fin], arr[i]
	return arr, i
}

func QuickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func main() {
	var slice = make([]int, 30)
	for i := 0; i < 30; i++ {
		slice[i] = rand.Intn(999)
	}
	slice2 := append([]int(nil), slice...)
	fmt.Println(slice)
	sort.Slice(slice, func(a, b int) bool { return slice[a] < slice[b] })
	fmt.Println(slice)
	fmt.Println(QuickSortStart(slice2))

}
