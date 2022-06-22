package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, result float64 = 0, 0, 0, 0
	fmt.Scanln(&a, &b, &c)

	if a == b && b == c {
		result = 10000 + (1000 * a)
	} else if a == b || b == c || a == c {
		if a == b {
			result = 1000 + (100 * a)
		} else if b == c {
			result = 1000 + (100 * b)
		} else {
			result = 1000 + (100 * c)
		}
	} else {
		result = math.Max(a, math.Max(b, c)) * 100

	}
	fmt.Println(int(result))
}
