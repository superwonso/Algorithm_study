package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var dp [][]int

func main() {
	var n, k int
	var result []int
	var t int
	dp = make([][]int, 31)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 31)
	}

	fmt.Fscan(r, &t)

	for i := 0; i < t; i++ {
		fmt.Scanln(&n, &k)
		result = append(result, Combination(k, n))
	}
	for i := 0; i < t; i++ {
		fmt.Println(result[i])
	}
}

func Combination(m, n int) int {
	if m == n || n == 0 {
		return 1
	}
	if dp[m][n] != 0 {
		return dp[m][n]
	}
	dp[m][n] = Combination(m-1, n) + Combination(m-1, n-1)
	return dp[m][n]
}
