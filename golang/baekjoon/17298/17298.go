package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)
	A := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanln(&A[i])
	}

	result := make([]int, N)
	for x := 0; x < N; x++ {
		for i := x; i < len(A)-1; i++ {
			result[x] = A[i]
			if A[i] < A[i+1] {
				result[x] = A[i]
			}
			if result[x] == A[i] {
				result[x] = -1
			}
		}
	}
	fmt.Println(result)
}
