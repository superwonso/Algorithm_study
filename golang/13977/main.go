package main

import (
	"fmt"
)

func main() {
	var m int
	var n, k int
	var result [100000]int
	fmt.Scanln(&m)
	for a := 0; a < m; a++ {
		fmt.Scanln("%d %d", &n, &k)
		result[a] = Combination(n, k) % 10007
	}
	for a := 0; a < m; a++ {
		fmt.Println(result[a])
	}
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
