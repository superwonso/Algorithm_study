package main

import "fmt"

func main() {
	var N, result int
	fmt.Scanln(&N)
	for i := 1; i <= N; i++ {
		if check_hansoo(i) == true {
			result++
		}
	}
	fmt.Println(result)
}

func check_hansoo(n int) bool {
	if n < 100 {
		return true
	} else if n >= 100 && n < 1000 {
		if (n/100)-((n/10)%10) == ((n/10)%10)-(n%10) {
			return true
		}
	} else {
		return false
	}
	return false
}
