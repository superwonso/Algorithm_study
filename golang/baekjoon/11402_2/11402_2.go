package main

import (
	"fmt"
)

func main() {
	var n, k, m int
	fmt.Scanln(&n, &k, &m)
	var result int = 1
	for n > 0 || k > 0 {
		result *= Combination(n%m, k%m, m) % m
		n /= m
		k /= m
	}
	fmt.Println(result)
}

func Combination(n, k, m int) int {
	var a, b int = 1, 1
	for i := n; i > n-k; i-- {
		a *= i % m
	}
	for i := k; i >= 2; i-- {
		b *= i % m
	}
	n = 1
	k = m - 2
	for k > 0 {
		if k%2 == 1 {
			n *= b % m
		}
		k >>= 1
		b *= b % m
	}
	return a * n % m
}
