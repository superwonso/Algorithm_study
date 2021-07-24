package main

import (
	"fmt"
)

func PrintCombination() {
	var n, k int
	fmt.Scanln(&n, &k)
	fmt.Println(Combination(n, k))
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
