package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	n1, err := big.NewInt(0).SetString("12345812838182388123812318231", 10)
	if !err {
		log.Fatal(err)
	}
	n2, err := big.NewInt(0).SetString("31412121412424124124124121211", 10)
	if !err {
		log.Fatal(err)
	}
	calculator := big.NewInt(0)
	fmt.Println("Sum: ", calculator.Add(n1, n2))
	fmt.Println("Sub: ", calculator.Sub(n1, n2))
	fmt.Println("Mul: ", calculator.Mul(n1, n2))
	fmt.Println("Div: ", calculator.Div(n1, n2))

}
