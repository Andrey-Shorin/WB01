package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Scan()
	var s []string = []string(strings.Fields(scaner.Text()))
	slices.Reverse(s)
	fmt.Println(strings.Join(s, " "))
}
