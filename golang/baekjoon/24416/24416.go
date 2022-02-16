//not solved
package main

import (
	"bufio"
	"fmt"
	"os"
)

var count_recur, count_dyn int = 0, 0

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()
	var n int
	fmt.Fscanf(r, "%d", &n)
	fib_recur(n)
	fib_dynamic(n)
	fmt.Fprintln(w, count_recur, count_dyn)
}

func fib_recur(n int) int {
	if n == 1 {
		count_recur = 1
		return 1
	} else if n == 2 {
		count_recur = 2
		return 1
	} else if n <= 1 {
		return 0
	}
	fib_recur_count()
	return fib_recur(n-1) + fib_recur(n-2)

}

func fib_recur_count() int {
	count_recur++
	return count_recur
}

func fib_dynamic(n int) int {
	var fib_dyn [40]int
	fib_dyn[0] = 0
	fib_dyn[1] = 1
	fib_dyn[2] = 1
	if n == 1 {
		count_dyn = 1
		return count_dyn
	} else if n == 2 {
		count_dyn = 2
		return count_dyn - 1
	}

	for count_dyn = 2; count_dyn <= n; count_dyn++ {
		fib_dyn[count_dyn] = fib_dyn[count_dyn-1] + fib_dyn[count_dyn-2]
	}
	return fib_dyn[n]
}
