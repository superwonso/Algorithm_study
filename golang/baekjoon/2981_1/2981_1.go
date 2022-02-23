//brute force -> timeout
package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	inp := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&inp[i])
	}
	tmp := make([][]int, min(inp), min(inp))
	for i := 0; i < n; i++ {
		for j := 1; j <= min(inp); j++ {
			tmp[i][j] = inp[i] % j
		}
	}
}

func min(input []int) int {
	min_value := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] < min_value {
			min_value = input[i]
		}
	}
	return min_value
}
