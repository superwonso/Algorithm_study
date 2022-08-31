package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)

func main() {
	var n, result int64
	fmt.Fscan(r, &n)
	var i int64 = 2
	result = n
	for i = 2; i*i <= n; i++ {
		if n%i == 0 {
			result -= result / i
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		result -= result / n
	}
	fmt.Println(result)
}
