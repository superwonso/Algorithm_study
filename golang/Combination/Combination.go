package main

import (
	"fmt"
)

func Return_Combination() int {
	var n, k, result int
	fmt.Scanln(&n, &k)
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
