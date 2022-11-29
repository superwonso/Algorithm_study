package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)

func main() {
	var a, b, c int = 0, 0, 0
	fmt.Fscan(r, &a, &b, &c)
	for i := 0; i < b; i++ {
		a = modulo_distributive(a, a, c)
	}
	fmt.Println(a)
}

/*func get_Digits(n int) int {
	var digit int
	for n > 0 {
		digit++
		n /= 10
	}
	return digit
}*/

func modulo_distributive(a, b, c int) int {
	return (a % c) * (b % c) % c
}
