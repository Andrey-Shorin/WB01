package main

import (
	"fmt"
	"math/rand"
)

func someFunc(n int) string {
	sliseofbyte := createHugeString(n)
	justbytes := make([]byte, 100)
	copy(justbytes, sliseofbyte[:100])
	return string(justbytes)
}
func main() {
	fmt.Println(someFunc((1 << 10)))

}

func createHugeString(n int) []byte {
	var str string = "abcdefghijklmnopqrstuvwxyz0123456789QWERTYUIOPLKJHGFDSAZXCVBNM"
	sliseofbyte := make([]byte, n, n)
	for i := 0; i < n; i++ {
		sliseofbyte[i] = str[rand.Intn(len(str))]
	}
	return sliseofbyte

}
