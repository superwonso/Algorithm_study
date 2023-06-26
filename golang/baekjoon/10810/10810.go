package main

import "fmt"

func main() {
	var N, M int
	fmt.Scanln(&N, &M)
	k := make([]int, N)
	for i := 0; i < M; i++ {
		var a, b, c int
		fmt.Scanln(&a, &b, &c)
		// from a to b, value is c
		for j := a; j <= b; j++ {
			k[j-1] = c
		}
	}
	for i := 0; i < N; i++ {
		fmt.Printf("%d ", k[i])
	}
}
