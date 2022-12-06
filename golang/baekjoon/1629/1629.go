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
	a = modulo_distributive(a, b, c)
	fmt.Println(a)
}

func modulo_distributive(a, b, c int) int {
	if b == 1 {
		return a % c
	} else if b%2 == 0 {
		tmp := modulo_distributive(a, b/2, c)
		return (tmp % c) * (tmp % c) % c
	}
	return modulo_distributive(a, b/2, c) * modulo_distributive(a, (b/2)+1, c) % c
}
