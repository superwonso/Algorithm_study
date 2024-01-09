package main

import (
	"bufio"
	"fmt"
	"os"
)

var count_recur, count_dyn int = 0, 0
var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	var n int
	fmt.Fscanf(r, "%d", &n)
	fib_recur(n)
	fib_dynamic(n)
	fmt.Fprintf(w, "%d %d", count_recur, count_dyn)
	defer w.Flush()
}

func fib_recur(n int) int {
	if n == 1 {
		count_recur++
		return 1
	} else if n == 2 {
		count_recur++
		return 1
	} else if n < 1 {
		return 0
	}
	return fib_recur(n-1) + fib_recur(n-2)
}

var fib_dyn [41]int

func fib_dynamic(n int) int {
	fib_dyn[0] = 0
	fib_dyn[1] = 1
	fib_dyn[2] = 1
	if n == 1 {
		count_dyn++
		return 1
	} else if n == 2 {
		count_dyn++
		return 1
	}
	for i := 3; i <= n; i++ {
		count_dyn++
		fib_dyn[i] = fib_dyn[i] + fib_dyn[i]
	}
	return fib_dyn[n]
}
