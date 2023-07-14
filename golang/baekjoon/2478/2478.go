package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	var arr [501]int
	var tmp [501]int
	var n, s, e int
	var visited [501]bool

	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}

	for i := 0; i < n; i++ {
		if arr[(n+i-1)%n]-arr[i] == 1 {
			visited[(n+i-1)%n], visited[i] = true, true
		}
		if arr[(i+1)%n] == 1 && arr[i] == n {
			visited[(n+i-1)%n], visited[i] = true, true
		}
	}

	s, e = -1, -1
	for i := 0; i < n; i++ {
		if visited[(n+i-1)%n] == false && visited[i] {
			s = i
		}
		if visited[i] && visited[(i+1)%n] == false {
			e = i
		}
	}
	if s == -1 {
		s, e = 0, n-1
	}
	sh := 0
	if s > e {
		sh, n = s, s
		for i := 0; i < n; i++ {
			tmp[i] = arr[i]
		}
		for i := 0; i < n; i++ {
			arr[i] = tmp[(n+1-sh)%n]
		}
		s, e = (s+sh)%n, (e+sh)%n
	}
	reverse(arr[:], s, e)
	for i := 0; i < n; i++ {
		if arr[i] == 1 {
			fmt.Fprintf(w, "%d\n", n-i)
			break
		}
	}
	fmt.Fprintf(w, "%d %d\n", s+1, e+1)
	//print sh? sh: N below
	fmt.Fprintf(w, "%d\n", sh)
	w.Flush()
}

// declare reverse function in c++
func reverse(arr []int, s, e int) {
	for i := 0; i < (e-s+1)/2; i++ {
		arr[s+i], arr[e-i] = arr[e-i], arr[s+i]
	}
}
