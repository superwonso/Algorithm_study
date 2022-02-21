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
	} else if check_prime(p) == false {
		var power int = 0
		power = a
		fmt.Println(power)
		for i := 1; i < p; i++ {
			power = power * a
		}
		fmt.Println(power)
		if power%p != a {
			return "no"
		} else {
			return "yes"
		}
	} else {
		return "fail"
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
