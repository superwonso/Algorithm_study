package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var create, dismantle int
	n := bufio.NewReader(os.Stdin)
	fmt.Fscan(n, &dismantle)

	for i := 0; i < dismantle; i++ {
		var digits = getcreate(i)
		if digits == dismantle {
			create = i
			break
		}
	}
	fmt.Println(create)
}

func getcreate(n int) (result int) {
	result = n
	for n != 0 {
		result += n % 10
		n /= 10
	}
	return result
}
