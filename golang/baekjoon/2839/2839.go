package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)

func main() {
	var N, v1, v2, result int = 0, 3, 5, 0
	fmt.Fscan(r, &N)
	for true {
		if N%v2 != 0 {
			if N < v1 {
				result = -1
				break
			}
			N -= v1
			result++
		} else {
			break
		}
	}
	if result != -1 {
		result += N / v2
	}
	fmt.Println(result)
}
