package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Tree struct {
	node *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	var depth int
	tr := &Tree{}

	fmt.Fscanln(r, &depth)
	var n int = 1

	for i := 0; i < depth; i++ {
		n *= 2
	}

	var arr = make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscanf(r, "%d", &arr[i])
	}
	fmt.Fprint(w, arr, "\n")

	for i := 0; i < n-1; i++ {
		tr.insert(arr[i])
		fmt.Fprintln(w, "arr :  arr[i]", arr[i])
		fmt.Fprintln(w, "tr.node.value: ", tr.node.value)
	}

	//tr = make_complete_binary_tree(tr, depth)

	// fmt.Fprintln(w, "tr.node.value: ", tr.node.value)
	// fmt.Fprintln(w, "tr.node.left.value: ", tr.node.left.value)

	for i := 0; i < depth; i++ {
		for j := 0; j < int(math.Pow(2, float64(i))); j++ {
			// if i == 0 {
			// 	fmt.Fprintf(w, "%d", tr.node.value)
			// } else {
			// 	fmt.Fprintf(w, "%d ", tr.node.value)
			// }
			printNode_in_level(tr.node, i+1)
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
}

func (t *Tree) insert(value int) *Tree {

	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.insert(value)
	}

	return t
}

func (n *Node) insert(value int) {
	if n.left == nil {
		n.left = &Node{value: value}
	} else {
		n.left.insert(value)
	}
	if n.right == nil {
		n.right = &Node{value: value}
	} else {
		n.right.insert(value)
	}
	return
}

func printNode_in_level(n *Node, level int) {
	if n == nil {
		return
	}
	if level == 1 {
		fmt.Fprintf(w, "%d ", n.value)
	}
	printNode_in_level(n.left, level-1)
	printNode_in_level(n.right, level-1)
	return
}

func make_complete_binary_tree(t *Tree, depth int) *Tree {
	if depth == 1 {
		return t
	}
	t.node.left = make_complete_binary_tree(t.node.left, depth-1)
	t.node.right = make_complete_binary_tree(t.node.right, depth-1)
	return t
}
