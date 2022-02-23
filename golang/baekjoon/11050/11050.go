package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k int
	fmt.Println(Return_Combination(n, k))
}

func Return_Combination(n, k int) int {
	var result int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &k)
	result = Combination(n, k)
	return result
}

func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return (n * factorial(n-1))
}

func Permutation(n, k int) int {
	return factorial(n) / factorial(n-k)
}

func Combination(n, k int) int {
	return Permutation(n, k) / factorial(k)
}
