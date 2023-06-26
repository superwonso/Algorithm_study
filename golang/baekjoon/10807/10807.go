package main

import "fmt"

func main() {
	var N int
	fmt.Scanln(&N)
	k := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &k[i])
	}
	var index int
	fmt.Scanln(&index)
	var count int
	for i := 0; i < N; i++ {
		if k[i] == index {
			count++
		}
	}
	fmt.Printf("%d", count)
}
