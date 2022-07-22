//Not Solved
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, res int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &N)
	A := make([]int, N)
	tmp := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &A[i])
	}
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			if A[i] > A[j] {
				tmp[i] = max(tmp[i], tmp[j]+1)
			}
		}
	}
	res = tmp[0]
	for i := 0; i < N; i++ {
		if res < tmp[i] {
			res = tmp[i]
		}
	}
	fmt.Println(res + 1)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
