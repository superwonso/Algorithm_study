package main

import (
	"fmt"
)

var s []string

func main() {
	var n int

	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		s = append(s, convert_to_binary(i))
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s ", s[i])
	}
}

func convert_to_binary(n int) string {
	if n == 0 {
		return "0"
	}
	convert_to_binary(n / 2)
	return "1"
}
