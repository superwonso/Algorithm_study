package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)

func main() {
	var X, t, result int
	fmt.Fscan(r, &X)
	fmt.Fscan(r, &t)
	for i := 0; i < t; i++ {
		var p, n int
		fmt.Fscan(r, &p, &n)
		result += p * n
	}
	if result == X {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
