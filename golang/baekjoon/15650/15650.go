package main

// Solved Not Yet
import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	result := make([]int, M)

	result[M-1] = rand.Intn(N-1) + 1
	for i := 0; i < M-1; i++ {
		result[i] = rand.Intn(N-1) + 1
		if check_overlap(result[i], result[i+1]) {
			i--
		}
	}
	sort.Sort(sort.IntSlice(result))
	for i := 0; i < M; i++ {
		fmt.Println(result[i])
	}
}

func check_overlap(a, b int) bool {
	if a == b {
		return true
	}
	return false
}
