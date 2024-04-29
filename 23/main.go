package main

import (
	"fmt"
)

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sl)
	sl, _ = deleteindex(sl, 2)
	fmt.Println(sl)
	sl, _ = deleteindex(sl, 2)
	fmt.Println(sl)
}

func deleteindex(sl []int, index int) ([]int, error) {
	if index < 0 || index >= len(sl) {
		return nil, fmt.Errorf("wrong index")
	}
	return append(sl[0:index], sl[index+1:]...), nil
}
