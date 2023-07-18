package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	var money int
	var coin = [6]int{500, 100, 50, 10, 5, 1}
	var cnt int = 0
	fmt.Fscanf(r, "%d\n", &money)
	var tmp = 1000 - money
	for i := 0; i < 6; i++ {
		cnt += tmp / coin[i]
		tmp %= coin[i]
	}
	fmt.Fprintf(w, "%d\n", cnt)
	w.Flush()
}
