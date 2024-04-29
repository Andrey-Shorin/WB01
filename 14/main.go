package main

import (
	"log"
	"reflect"
)

// Разработать программу, которая в рантайме способна
// определить тип переменной: int, string, bool, channel из
// переменной типа interface{}.

func main() {
	log.Println(checkType(nil))
	log.Println(checkType("ytterby"))
	log.Println(checkType(199))
	log.Println(checkType([]int{1, 2, 3}))
	log.Println(checkType(make(chan interface{})))

}
func checkType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}
