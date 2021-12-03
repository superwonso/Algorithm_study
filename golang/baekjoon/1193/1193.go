package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)
	result := make([]int, n+1)
	count := 2
	for i := 1; i < len(result); i++ {
		if ((i*(i+1))/2)-i >= i {
			count++
		}
		if n == 1 {
			result[i] = count
		} else if n%2 == 0 && n > 2 {
			result[i] = count
		} else if n%2 == 1 && n > 2 {

		}
	}

}
