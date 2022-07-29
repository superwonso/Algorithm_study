package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	var n, m int
	fmt.Fscanln(r, &n, &m)
	var visited = make([]bool, n+1)
	var result = make([]int, m+1)
	sequence(visited, result, 0, n, m)
	defer w.Flush()
}

func sequence(visited []bool, result []int, index, n, m int) {
	if index == m {
		for i := 0; i < m; i++ {
			fmt.Fprintf(w, "%d ", result[i])
		}
		fmt.Fprint(w, "\n")
		return
	}
	for i := 1; i < n+1; i++ {
		if visited[i] || index > 0 && result[index-1] > i {
			continue
		}
		visited[i] = true
		result[index] = i
		sequence(visited, result, index+1, n, m)
		visited[i] = false
	}
}
