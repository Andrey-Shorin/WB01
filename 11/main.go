package main

import "fmt"

func main() {
	setOne := []int{99, 0, 37, -17, 5, 6}
	setTwo := []int{5, 0, -99, -14, 8, 3}
	fmt.Println(Intersect(setOne, setTwo))
}

func Intersect(firstSet []int, secondSet []int) []int {

	hash := make(map[int]bool)
	for _, i := range firstSet {
		hash[i] = true
	}
	var result []int
	for _, j := range secondSet {
		if _, ok := hash[j]; ok {
			result = append(result, j)
		}
	}
	return result
}
