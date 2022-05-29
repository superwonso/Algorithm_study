// O(n^2) , Fail
package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)
	A := make([]int, N)
	result := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}
	result = A
	for x := 0; x < N; x++ {
		for y := x; y < N; y++ {
			if result[x] < A[y] {
				result[x] = A[y]
				fmt.Scanln(result[x])
			}
		}
		if result[x] == A[x] {
			result[x] = -1
		}
	}
	fmt.Println(result)
}
