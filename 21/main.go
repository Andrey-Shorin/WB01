package main

import (
	"fmt"
)

// Реализовать паттерн «адаптер» на любом примере.

type Namer interface {
	Name() string
}

func Greater(n Namer) {
	fmt.Println("Hello", n.Name())
}

type Person struct {
	login string
}

func (p Person) Login() string {
	return p.login
}

type Pet struct {
	nickname string
}

func (p Pet) Nick() string {
	return p.nickname
}

type NamePersonAdapter struct {
	Person
}

func (n NamePersonAdapter) Name() string {
	return n.Login()
}

func main() {
	p := Person{login: "Vasya"}

	Greater(NamePersonAdapter{p})
}
