package main

import "fmt"

func main() {
	var N int
	var charr string
	charr = "long "
	fmt.Scanln(&N)
	for i := 0; i < N/4; i++ {
		fmt.Printf("%s", charr)
	}
	fmt.Printf("int")
}
