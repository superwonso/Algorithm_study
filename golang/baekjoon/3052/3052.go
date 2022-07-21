package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := make([]int, 10)
	result := make([]int, 42)
	count := 0
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		fmt.Fscan(r, &input[i])
	}
	for i := 0; i < 10; i++ {
		input[i] = input[i] % 42
	}
	for i := 0; i < 10; i++ {
		result[input[i]]++
	}
	for i := 0; i < 42; i++ {
		if result[i] != 0 {
			count++
		}
	}
	fmt.Println(count)
}
