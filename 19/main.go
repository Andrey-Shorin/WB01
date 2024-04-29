package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Scan()
	var s []rune = []rune(scaner.Text())
	slices.Reverse(s)
	fmt.Println(string(s))
}
