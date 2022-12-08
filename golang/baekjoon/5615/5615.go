package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	var n, cnt int64 = 0, 0

	fmt.Fscan(reader, &n)

	tmp2 := big.NewInt(1)
	var i int64 = 0

	for i = 1; i <= n; i++ {
		var s string
		fmt.Fscan(reader, &s)

		tmp, _ := new(big.Int).SetString(s, 10)
		tmp.Add(tmp, tmp)
		tmp.Add(tmp, tmp2)

		// about ProbablyPrime function... https://golang.org/pkg/math/big/#Int.ProbablyPrime
		// ProbablyPrime is 100% accurate for inputs less than 2^64, for larger inputs it is accurate (1 - 1/4^n) when set function as x.ProbablyPrime(n).
		// If n==0, Up to Go 1.8, ProbablyPrime(0) is apply only a Baillie-PSW primality test.
		if tmp.ProbablyPrime(0) {
			cnt++
		}

	}
	fmt.Println(cnt)
}
