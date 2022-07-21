package main

import "fmt"

func main() {
	selfNumbers := self_Number(10000)
	for i := 1; i < len(selfNumbers); i++ {
		if selfNumbers[i] == 0 {
			fmt.Println(i)
		}
	}
}

func self_Number(n int) []int {
	check_self_number := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		check_self_number[i] = 0
	}

	for i := 0; i < n+1; i++ {
		var sum = i
		var number = i
		for j := number; j != 0; j /= 10 {
			sum += j % 10
		}
		if sum <= n {
			check_self_number[sum] = 1
		}
	}

	return check_self_number
}
