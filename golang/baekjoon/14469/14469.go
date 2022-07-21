package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, res int
	fmt.Scanln(&N)
	cow := make([][]int, N)
	for i := 0; i < N; i++ {
		cow[i] = make([]int, 2)
	}
	for i := 0; i < N; i++ {
		fmt.Scanln(&cow[i][0], &cow[i][1])
	}
	sort.Slice(cow, func(i, j int) bool { return cow[i][0] < cow[j][0] })
	res = cow[0][0] + cow[0][1]
	for i := 1; i < N; i++ {
		if res <= cow[i][0] {
			res = cow[i][0] + cow[i][1]
		} else {
			res += cow[i][1]
		}
	}
	fmt.Println(res)
}
