package main

// Solved Not Yet

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, idx, temp int = 0, 0, 0
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &N)
	tmp := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscan(r, &temp)
		if idx == 0 {
			idx++
			tmp[idx] = temp
		} else {
			if tmp[idx-1] < temp {
				idx++
				tmp[idx] = temp
			} else {
				tmp[lower_bound(tmp, i, temp)] = temp
			}
		}
	}
	fmt.Println(idx)
}

func lower_bound(arr []int, n int, key int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r) / 2
		if arr[mid] < n {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
