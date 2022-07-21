package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scanln(&A)
	fmt.Scanln(&B)
	fmt.Scanln(&C)
	tmp := A * B * C
	calc := make([]int, 10)
	for true {
		calc[tmp%10]++
		tmp /= 10
		if tmp == 0 {
			break
		}
	}
	for i := 0; i < len(calc); i++ {
		fmt.Println(calc[i])
	}
}
