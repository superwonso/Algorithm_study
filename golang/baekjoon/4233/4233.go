package main

import (
	"fmt"
)

func main() {
	var a, p int
	var repeat bool = true
	for repeat == true {
		fmt.Scanln(&p, &a)
		if a == 0 && p == 0 {
			repeat = false
			break
		} else {
			fmt.Println(prove(a, p))
		}

	}
}

func prove(a, p int) string {
	if check_prime(p) == true {
		return "no"
	} else {
		if mod_pow(a, p, p) == a {
			return "yes"
		} else {
			return "no"
		}
	}
}

func check_prime(p int) bool {
	if p == 2 {
		return true
	}
	if p%2 == 0 {
		return false
	}
	for i := 3; i*i <= p; i += 2 {
		if p%i == 0 {
			return false
		}
	}
	return true
}

//Divide and Conquer
func mod_pow(a, p, m int) int {
	var result int = 1
	for p > 0 {
		if p%2 == 1 {
			result = result * a % m
		}
		a = a * a % m
		p = p / 2
	}
	return result
}
