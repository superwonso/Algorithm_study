package main

import (
	"bufio"
	"fmt"
	"os"
)

var top int = -1
var stack [100001]int

func main() {
	var k, result, tmp int
	fmt.Scanln(&k)
	for i := 0; i < k; i++ {
		r := bufio.NewReader(os.Stdin)
		fmt.Fscan(r, &tmp)
		if tmp == 0 {
			pop()
		} else {
			push(tmp)
		}
	}
	for i := 0; i <= top; i++ {
		result += stack[i]
	}

	fmt.Printf("%d", result)
}

func push(input int) {
	top++
	stack[top] = input
}

func pop() {
	top--
}
