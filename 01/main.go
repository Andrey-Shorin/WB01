package main

import (
	"log"
)

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание(композиция) методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	Name string
}

func (h *Human) Say() {
	log.Println("I'm ", h.Name)
}

type Action struct {
	Human
}

func main() {
	a := Action{Human{Name: "Tom"}}
	a.Say()
}
