package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var numbers = make([]int, 9)
	r := bufio.NewReader(os.Stdin)
	var max = 0
	var maxI = 0
	for i := range numbers {
		fmt.Fscan(r, &numbers[i])
		if numbers[i] > max {
			max = numbers[i]
			maxI = i
		}
	}

	fmt.Println(max)
	fmt.Println(maxI + 1)
}
