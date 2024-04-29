package main

import (
	"fmt"
	"strings"
)

func checstr(str string) bool {
	mymap := make(map[string]bool)
	for _, i := range str {
		if mymap[strings.ToUpper(string(i))] {
			return false
		}
		mymap[strings.ToUpper(string(i))] = true
	}
	return true
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(checstr(str))
}
