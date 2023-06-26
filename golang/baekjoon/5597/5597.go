package main

import "fmt"

func main() {
	var k [30]bool
	for i := 0; i < 30; i++ {
		k[i] = false
	}
	for i := 0; i < 28; i++ {
		var N int
		fmt.Scanln(&N)
		k[N-1] = true
	}
	for i := 0; i < 30; i++ {
		if k[i] == false {
			fmt.Printf("%d\n", i+1)
		}
	}
}
