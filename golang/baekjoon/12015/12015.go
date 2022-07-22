package main

import (
	"bufio"
	"fmt"
	"os"
)

var lis []int
var Arr [1000001]int

func main() {
	var N, l_idx int = 0, 0
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &Arr[i])
	}
	lis = push_back(Arr[0])
	for i := 1; i < N; i++ {
		if lis[l_idx] < Arr[i] {
			lis = push_back(Arr[i])
			l_idx++
		} else {
			var idx int = lower_bound(Arr[i])
			lis[idx] = Arr[i]
		}
	}
	fmt.Println(len(lis))
}

func lower_bound(n int) int {
	var l int = 0
	var r int = len(lis) - 1
	for l < r {
		var mid int = (l + r) / 2
		if lis[mid] >= n {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func push_back(n int) []int {
	lis = append(lis, n)
	return lis
}
