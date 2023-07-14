package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	var n int
	fmt.Fscan(r, &n)
	defer w.Flush()
	fmt.Fprintln(w, (1<<n)-1)
	hanoi(n, 1, 3, 2)
}

func hanoi(n, from, to, via int) {
	if n == 1 {
		fmt.Fprintf(w, "%d %d\n", from, to)
		return
	}
	hanoi(n-1, from, via, to)
	move(from, to)
	hanoi(n-1, via, to, from)
}

func move(from, to int) {
	fmt.Fprintf(w, "%d %d\n", from, to)
}
