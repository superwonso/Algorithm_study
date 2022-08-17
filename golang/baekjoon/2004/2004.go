package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)

func main() {

	var n, m, count, tC, fC int
	fmt.Fscanln(r, &n, &m)

	tC = getTwoCount(n) - getTwoCount(n-m) - getTwoCount(m)
	fC = getFiveCount(n) - getFiveCount(n-m) - getFiveCount(m)

	if tC > fC {
		count = fC
	} else {
		count = tC
	}

	fmt.Println(count)
}

func getTwoCount(num int) (tC int) {
	for i := 2; i <= num; i *= 2 {
		tC += num / i
	}

	return tC
}

func getFiveCount(num int) (fC int) {
	for i := 5; i <= num; i *= 5 {
		fC += num / i
	}

	return fC
}
