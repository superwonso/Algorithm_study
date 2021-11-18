package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)
	tc := make([]int, t)
	a := make([]int, t)
	b := make([]int, t)
	c := make([]int, t)
	var i int = 0
	for i = 0; i < t; i++ {
		fmt.Scanln(&a[i], &b[i])
		tc[i] = i + 1
		c[i] = a[i] + b[i]
		if c[i] >= 24 {
			c[i] -= 24
		} else if c[i] == 48 {
			c[i] = 0
		}
	}
	for i = 0; i < t; i++ {
		fmt.Printf("#%d %d\n", tc[i], c[i])
	}
}
