package main

import "fmt"

const N int = 8

var kit [N]int
var chk [N]bool
var n, K int

func runi(w int) int {
	left := false
	sum := 0
	for i := 0; i < n; i++ {
		if !chk[i] {
			left = true
			nw := w + kit[i] - K
			if nw >= 0 {
				chk[i] = true
				sum += runi(nw)
				chk[i] = false
			}
		}
	}
	if !left {
		return 1
	} else {
		return sum
	}
}

func main() {
	fmt.Scanf("%d %d", &n, &K)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &kit[i])
	}
	fmt.Println(runi(0))
}
