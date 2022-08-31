package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)

func main() {
	var tc int
	fmt.Fscan(r, &tc)
	for i := 0; i < tc; i++ {
		var num int
		fmt.Fscan(r, &num)
		var s string
		fmt.Fscan(r, &s)
		for j := 0; j < len(s); j++ {
			for k := 0; k < num; k++ {
				fmt.Printf("%c", s[j])
			}
		}
		fmt.Println()
	}
}
