package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var res []int
	for n > 1 {
		divisior := pollard_rho(n)
		res = append(res, divisior)
		n /= divisior
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d\n", res[i])
	}
}

func power(x, y, p int) int {
	res := 1
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

func gcd(a, b int) int {
	var r int
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		r = a % b
		a = b
		b = r
	}
	return a
}

func miller_rabin(n, a int) bool {
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

func isPrime(n int) bool {
	var test_set = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
	if n == 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 0; i < len(test_set); i++ {
		if n == i {
			return true
		} else if miller_rabin(n, i) {
			return false
		}
	}
	return true
}

func pollard_rho(n int) int {
	if isPrime(n) {
		return n
	}
	if n == 1 {
		return 1
	}
	if n%2 == 0 {
		return 2
	}
	x := rand.Intn(n-2) + 2
	y := x
	c := rand.Intn(n-1) + 1
	d := 1
	for d == 1 {
		x = ((int((math.Pow(float64(x), 2))) % n) + c + n) % n
		y = ((int((math.Pow(float64(y), 2))) % n) + c + n) % n
		y = ((int((math.Pow(float64(y), 2))) % n) + c + n) % n
		d = gcd(int(math.Abs(float64(x-y))), n)
		if d == n {
			return pollard_rho(n)
		}
	}
	if isPrime(d) {
		return d
	} else {
		return pollard_rho(d)
	}
}
