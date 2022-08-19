package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)

func main() {
	var a, b, v int
	fmt.Fscan(r, &a, &b, &v)
	fmt.Println((v-b-1)/(a-b) + 1)
}
