package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)

func main() {
	var status int = 1
	n := make([]int, 8)
	for i := range n {
		fmt.Fscan(r, &n[i])
	}
	for i := 0; i < len(n)-1; i++ {
		if n[i]-n[i+1] == -1 {
			status = 1
		} else if n[i]-n[i+1] == 1 {
			status = 0
		} else {
			status = -1
			break
		}
	}
	if status == 1 {
		fmt.Println("ascending")
	} else if status == 0 {
		fmt.Println("descending")
	} else if status == -1 {
		fmt.Println("mixed")
	}
}
