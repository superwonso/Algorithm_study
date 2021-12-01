package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var n string
		fmt.Fscanln(reader, &n)
		fmt.Fprintln(writer, check_valid_parenthetis_string(n))
	}
}

func check_valid_parenthetis_string(n string) string {
	var stack []rune
	for _, c := range n {
		if string(c) == "(" {
			stack = append(stack, c)
		} else if string(c) == ")" {
			if len(stack) == 0 {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}
