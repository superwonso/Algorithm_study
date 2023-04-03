package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println((1 << n) - 1)
	hanoi(n, "1", "3", "2")
}

func hanoi(n int, from, to, via string) {
	if n == 1 {
		fmt.Println(from, to)
	} else {
		hanoi(n-1, from, via, to)
		fmt.Printf("%s %s\n", from, to)
		hanoi(n-1, via, to, from)
	}
}
