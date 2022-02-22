package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k, m int
	fmt.Scanln(&n, &k, &m)
	var converted_n []int
	var converted_k []int
	converted_n = change_digits(n, m)
	converted_k = change_digits(k, m)
	var result int = 0
	for i := 0; i < len(converted_k); i++ {
		result += Combination(converted_n[i], converted_k[i])
		if converted_n[i] == 0 && converted_k[i] == 0 {
			break
		}
	}
	fmt.Println(result % m)
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

func change_digits(n, m int) []int {
	result := make([]int, n)
	var i int = 0
	for n > 0 {
		result[i] = n % m
		n = n / m
		i++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result
}
