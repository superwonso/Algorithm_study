package main

import (
	"bufio"
	"fmt"
	"os"
)

var Reader = bufio.NewReader(os.Stdin)

func main() {
	var k, q, r, b, n, p int
	fmt.Fscan(Reader, &k, &q, &r, &b, &n, &p)
	fmt.Println(1-k, 1-q, 2-r, 2-b, 2-n, 8-p)
}
