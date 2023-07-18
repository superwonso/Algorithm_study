package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	var n int
	fmt.Fscanln(r, &n)
	tree := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d", &tree[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(tree)))
	for i := 0; i < n; i++ {
		tree[i] += i + 1
	}
	sort.Sort(sort.Reverse(sort.IntSlice(tree)))
	fmt.Fprintln(w, tree[0]+1)
	w.Flush()
}
