package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	var N, M int
	var str string
	var result bool = false
	fmt.Fscan(r, &N, &M)
	fmt.Fscan(r, &str)
	for i := 0; i < M; i++ {
		result = false
		var postit string
		fmt.Fscan(r, &postit)
		// fmt.Fprintf(w, "postit: %s\n", postit)
		var k int = 0
		for j := 0; j < len(postit); j++ {
			if str[k] == postit[j] {
				k++
			}
			if k == N {
				result = true
				break
			}
		}
		if result == true {
			fmt.Fprintf(w, "true\n")
		} else {
			fmt.Fprintf(w, "false\n")
		}
	}
	w.Flush()
}
