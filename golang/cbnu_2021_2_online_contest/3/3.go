package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)
	a := make([]int, t)
	b := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Scanln(&a[i], &b[i])
	}

	for i := 0; i < t; i++ {
		if a[i] == 1 && b[i] == 2 {
			fmt.Printf("#%d B\n", i+1)
		} else if a[i] == 1 && b[i] == 3 {
			fmt.Printf("#%d A\n", i+1)
		} else if a[i] == 2 && b[i] == 1 {
			fmt.Printf("#%d A\n", i+1)
		} else if a[i] == 2 && b[i] == 3 {
			fmt.Printf("#%d B\n", i+1)
		} else if a[i] == 3 && b[i] == 1 {
			fmt.Printf("#%d B\n", i+1)
		} else if a[i] == 3 && b[i] == 2 {
			fmt.Printf("#%d A\n", i+1)
		}
	}

}
