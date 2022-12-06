package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)

func main() {
	var n, cnt int64 = 0, 0
	fmt.Fscan(r, &n)
	var i int64 = 0
	for i = 1; i <= n; i++ {
		var s int64
		fmt.Fscan(r, &s)
		if isPrime((2 * s) + 1) {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func power(x, y, p int64) int64 {
	var res int64 = 1
	//x = x % p
	for y != 0 {
		if y&1 == 1 {
			res = (res * x) % p
		}
		y >>= 1
		x = (x * x) % p
	}
	return res
}

func miller_rabin(n, a int64) bool {
	if a%n == 0 {
		return true
	}
	var d int64 = n - 1
	for d%2 == 0 {
		if power(a, d, n) == n-1 {
			return true
		}
		d /= 2
	}
	var temp int64 = power(a, d, n)
	return temp == 1 || temp == n-1
}

var test_set = []int64{2, 7, 61}

func isPrime(n int64) bool {
	for i := 0; i < len(test_set); i++ {
		if !miller_rabin(n, test_set[i]) {
			return false
		}
	}
	return true
}
