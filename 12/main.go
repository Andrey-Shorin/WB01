package main

import "fmt"

func main() {
	var str = []string{"cat", "dog", "cat", "cat", "tree"}
	set := make(map[string]struct{})
	for _, i := range str {
		set[i] = struct{}{}
	}
	fmt.Println(set)
}
