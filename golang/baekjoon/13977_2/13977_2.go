// Out of Bound(memory)
package main

import "fmt"

func binominal(n int, k int, dpTable *[][]int) int {
	if k == n || k == 0 || n == 1 {
		return 1
	}

	if (*dpTable)[n][k] > 0 {
		return (*dpTable)[n][k] % 1000000007
	}

	(*dpTable)[n][k] = binominal(n-1, k-1, dpTable) + binominal(n-1, k, dpTable)
	return (*dpTable)[n][k] % 1000000007
}

func main() {
	var N, K int
	MAX := 400000
	dpTable := make([][]int, 0)
	for i := 0; i < MAX; i++ {
		dpRow := make([]int, MAX)
		dpTable = append(dpTable, dpRow)
	}

	fmt.Scanf("%d %d", &N, &K)
	fmt.Printf("%d\n", binominal(N, K, &dpTable))
}
