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

		if tmp.ProbablyPrime(0) {
			cnt++
		}

	}
	fmt.Println(cnt)
}
