package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Tree struct {
	le  *Tree
	ri  *Tree
	val int
}

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	var depth int
	var tr Tree
	fmt.Fscanln(r, &depth)
	var n int = 1
	for i := 0; i < depth; i++ {
		n *= 2
	}
	var arr = make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscanf(r, "%d", &arr[i])
	}
	tr = *insert(&tr, arr[0])
	for i := 1; i < n-1; i++ {
		insert(&tr, arr[i])
	}
	tr = *make_complete_binary_tree(&tr, depth-1)
	for i := 0; i < depth; i++ {
		for j := 0; j < int(math.Pow(2, float64(i))); j++ {
			if i == 0 {
				fmt.Fprintf(w, "%d\n", tr.val)
			} else {
				print_tree(&tr, i)
				fmt.Fprintf(w, " ")
			}
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, nil, v}
	}
	if v < t.val {
		t.le = insert(t.le, v)
	} else {
		t.ri = insert(t.ri, v)
	}
	return t
}

func make_complete_binary_tree(t *Tree, k int) *Tree {
	if k == 0 {
		return t
	}
	if t.le != nil {
		t.le = make_complete_binary_tree(t.le, k-1)
	}
	if t.ri != nil {
		t.ri = make_complete_binary_tree(t.ri, k-1)
	}
	return t
}

// Print tree given level
// example... 1 6 4 3 5 2 7 -> 3; 6 2; 1 4 5 7;
func print_tree(t *Tree, k int) {
	if k == 0 {
		fmt.Fprintf(w, "%d ", t.val)
		return
	}
	if t.le != nil {
		print_tree(t.le, k-1)
	}
	if t.ri != nil {
		print_tree(t.ri, k-1)
	}
}
