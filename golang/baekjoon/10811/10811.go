package main

import "fmt"

func main() {
	var N, M int
	fmt.Scanln(&N, &M)
	k := make([]int, N)
	for i := 0; i < N; i++ {
		k[i] = i + 1
	}
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scanln(&a, &b)
		for j := 0; j <= (b-a)/2; j++ {
			k[j+a-1], k[b-j-1] = swap(k[j+a-1], k[b-j-1])
		}
	}
	for i := 0; i < N; i++ {
		fmt.Printf("%d ", k[i])
	}
}

func swap(a, b int) (int, int) {
	return b, a
}
