package main

import "fmt"

func main() {
	var m, n int64
	fmt.Scanln(&n, &m)
	result := n / m
	fmt.Println(result)
	fmt.Println(n % m)
}
