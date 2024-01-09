package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)
	var arr []int = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Printf("%d", solution_dp(arr, M))
}

func solution(arr []int, M int) int {
	var count int
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	for i := 0; i < len(arr); i++ {
		count += M / arr[i]
		M = M % arr[i]
	}
	if M != 0 {
		return -1
	}
	return count
}

func solution_dp(arr []int, M int) int {
	var dp []int = make([]int, M+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = 10001
	}
	for i := 0; i < len(arr); i++ {
		for j := arr[i]; j < len(dp); j++ {
			dp[j] = min(dp[j], dp[j-arr[i]]+1)
		}
	}
	if dp[M] == 10001 {
		return -1
	}
	return dp[M]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
