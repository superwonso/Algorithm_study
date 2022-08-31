// Solved Not Yet
// Overflow
package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func main() {
	var n int64
	fmt.Scanf("%d", &n)
	var res []int64
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

func gcd(a, b int64) int64 {
	var r int64
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
	var test_set = []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
	if n == 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	var i int64 = 0
	for i = 0; i < int64(len(test_set)); i++ {
		if n == i {
			return true
		} else if miller_rabin(n, i) {
			return false
		}
	}
	return true
}

func pollard_rho(n int64) int64 {
	if isPrime(n) {
		return n
	}
	if n == 1 {
		return 1
	}
	if n%2 == 0 {
		return 2
	}
	x := int64(rand.Intn(int(n-2))) + 2
	y := x
	c := int64(rand.Intn(int(n-1))) + 1
	var d int64 = 1
	for d == 1 {
		x = ((int64((math.Pow(float64(x), 2))) % n) + c + n) % n
		y = ((int64((math.Pow(float64(y), 2))) % n) + c + n) % n
		y = ((int64((math.Pow(float64(y), 2))) % n) + c + n) % n
		d = gcd(int64(math.Abs(float64(x-y))), n)
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
