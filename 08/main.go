package main

import (
	"errors"
	"fmt"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
func main() {
	var x int64 = 55
	var index1 int = 1
	fmt.Println(setone(x, index1))
	var index2 int = 1
	fmt.Println(setzero(x, index2))

}

func setone(number int64, index int) (int64, error) {
	if index >= 0 {
		return number | (1 << index), nil
	}
	return number, errors.New("index < 0")
}
func setzero(number int64, index int) (int64, error) {
	if index >= 0 {

		return number & ^(1 << index), nil
	}
	return number, errors.New("index < 0")
}
