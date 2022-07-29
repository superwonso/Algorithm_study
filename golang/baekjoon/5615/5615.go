package main

import "fmt"

func main() {
	var n, cnt int64 = 0, 0
	fmt.Scan(&n)
	for n > 0 {
		n--
		var s int64
		fmt.Scan(&s)
		if isPrime(2*s + 1) {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func power(x, y, p int64) int64 {
	var res int64 = 1
	x = x % p
	for y > 0 {
		if y&1 == 1 {
			res = (res * x) % p
		}
		y >>= 1
		x = (x * x) % p
	}
	return res
}

func miller_rabin(n, a int64) bool {
	r := 0
	d := n - 1
	for d%2 == 0 {
		r += 1
		d /= 2
	}
	x := power(a, d, n)
	if x == 1 || x == n-1 {
		return true
	}
	for i := 0; i < r-1; i++ {
		x = power(x, 2, n)
		if x == n-1 {
			return true
		}
	}
	return false
}

func isPrime(n int64) bool {
	var test_set = [5]int64{2, 3, 5, 7, 11}
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 0; i < len(test_set); i++ {
		var a int64 = test_set[i]
		if n == a {
			return true
		}
		if !miller_rabin(n, a) {
			return false
		}
	}
	return true
}
