package main

import "fmt"

func main() {
	var N, res int
	A := make([]int, N+1)
	tmp := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Scanf("%d", &A[i])
		tmp[i] = 1
	}
	for i := 1; i <= N; i++ {
		for j := 1; j < i; j++ {
			if A[i] > A[j] {
				max := func(n ...int) int {
					var x, y int
					if x > y {
						return x
					} else {
						return y
					}
				}
				tmp[i] = max(tmp[i], tmp[j]+1)
			}
		}
	}
	res = tmp[0]
	for i := 1; i <= N; i++ {
		if res < tmp[i] {
			res = tmp[i]
		}
	}
	fmt.Println(res)
}
