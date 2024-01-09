package main

import "fmt"

var dp [][]int

func main() {
	var m int
	var k int
	var results float64
	fmt.Scan(&m)
	colors := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&colors[i])
	}
	fmt.Scan(&k)

	var all int
	for i := 0; i < m; i++ {
		all += colors[i]
	}

	dp = make([][]int, all+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, all+1)
	}

	all = Combination(all, k)

	var each int
	for i := 0; i < m; i++ {
		if colors[i] >= k {
			each += Combination(colors[i], k)
		}
	}

	results = float64(each) / float64(all)
	fmt.Printf("%.10f", results)
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
