package main

// Not working yet

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var factorial [][]int // factorial[m][n] = (n!)%m
var inverse [][]int   // inverse[m][n] = (inverse of n!) %m
var p_adic_factorial [][]int
var p_adic_inverse [][]int
var n_list []int = []int{27, 11, 13, 37}

var r = bufio.NewReader(os.Stdin)

func main() {
	var t int
	fmt.Fscan(r, &t)
	for i := 0; i < t; i++ {
		var n, k int
		fmt.Fscan(r, &n, &k)
		a_list := Map3([]int{select_mod_prime(n, k, 3), select_mod_prime(n, k, 11), select_mod_prime(n, k, 13), select_mod_prime(n, k, 37)}, n_list, []int{3, 1, 1, 1}, func(a, b, c int) int { return a * int(math.Pow(float64(b), float64(c))) })
		fmt.Println(get_crt_root(a_list, n_list))
	}

}

func inverse_mod(a, b int) int {
	b0 := b
	x0, x1 := 0, 1
	if b == 1 {
		return 1
	}
	for a > 1 {
		q := a / b
		a, b = b, a
		x0, x1 = x1-q*x0, x0
	}
	if x1 < 0 {
		x1 += b0
	}
	return x1
}

func select_mod_prime(n, m, p int) int {
	gFactorial := func(num int) int {
		if !in(p, factorial) {
			factorial[p] = []int{1, 1}
		}
		for idx := 2; idx <= num; idx++ {
			factorial[p][idx] = factorial[p][idx-1] * idx % p
		}
		return factorial[p][num]
	}
	gInverse := func(num int) int {
		if !in(p, inverse) {
			inverse[p] = []int{1, 1}
		}
		for idx := 2; idx <= num; idx++ {
			inverse[p][idx] = inverse[p][p%idx] * (p - p/idx) % p
		}
		return inverse[p][num]
	}
	c_prod := 1
	for n != 0 || m != 0 {
		n_digit := n % p
		n /= p
		m_digit := m % p
		m /= p
		if n_digit < m_digit {
			c_prod = 0
			break
		} else {
			c_prod *= gFactorial(n_digit) * gInverse(m_digit) * gInverse(n_digit-m_digit)
			c_prod %= p
		}
	}
	return c_prod
}

func select_mod_prime_power(n, m, p, q int) int {
	modu := int(math.Pow(float64(p), float64(q)))

	var n_ex, m_ex, r_ex []int
	var d int
	expand_by_p := func(n, m, r int) ([]int, []int, []int, int) {
		d = 0
		d_min := q - 1
		for n > 0 || m > 0 || r > 0 || d <= d_min {
			n_ex = append(n_ex, n%p)
			n /= p
			m_ex = append(m_ex, m%p)
			m /= p
			r_ex = append(r_ex, r%p)
			r /= p
			d += 1
		}
		return n_ex, m_ex, r_ex, d - 1
	}

	var n_lpr, m_lpr, r_lpr []int
	least_positive_residue := func(n, m, r, d int) ([]int, []int, []int) {
		for i := 0; i < d+1; i++ {
			n_lpr = append(n_lpr, n%modu)
			n /= p
			m_lpr = append(m_lpr, m%modu)
			m /= p
			r_lpr = append(r_lpr, r%modu)
			r /= p
		}
		return n_lpr, m_lpr, r_lpr
	}

	var e0, eq1 int
	carry_out := func(m_ex, r_ex []int, d int) (int, int) {
		var has_carry []int
		prev_carry := 0
		for idx := 0; idx < d+1; idx++ {
			value := m_ex[idx] + r_ex[idx] + prev_carry
			if value >= p {
				has_carry[idx] = 1
				prev_carry = 1
			} else {
				prev_carry = 0
			}
		}

		for i := q - 1; i < len(has_carry); i++ {
			eq1 += +has_carry[i]
		}
		for i := 0; i < q-1; i++ {
			e0 += has_carry[i]
		}
		e0 += eq1
		return e0, eq1
	}

	get_p_adic_factorial := func(num int) int {
		if !in(p, p_adic_factorial) {
			p_adic_factorial[p] = []int{1, 1}
		}
		begin := len(p_adic_factorial[p])
		for idx := begin; idx < num+1; idx++ {
			var tmpr int
			if idx%p == 0 {
				tmpr = 1
			} else {
				tmpr = idx
			}
			p_adic_factorial[p][idx] = p_adic_factorial[p][idx-1] * tmpr % modu
		}
		return p_adic_factorial[p][num]
	}

	get_p_adic_inverse := func(num int) int {
		if !in(p, p_adic_inverse) {
			p_adic_inverse[p] = []int{1, 1}
		}
		begin := len(p_adic_inverse[p])
		for idx := begin; idx < num+1; idx++ {
			p_adic_factorial[p][idx] = inverse_mod(get_p_adic_factorial(idx), modu)
		}
		return p_adic_inverse[p][num]
	}

	r := n - m
	_, m_ex, r_ex, d = expand_by_p(n, m, r)
	n_lpr, m_lpr, r_lpr = least_positive_residue(n, m, r, d)
	e0, eq1 = carry_out(m_ex, r_ex, d)

	n_factorial := Map(n_lpr, get_p_adic_factorial)
	m_inverse := Map(m_lpr, get_p_adic_inverse)
	r_inverse := Map(r_lpr, get_p_adic_inverse)

	p_adic_choose := Map3(n_factorial, m_inverse, r_inverse, func(a, b, c int) int { return a * b * c % modu })
	choose_prod := reduce(Mul, p_adic_choose) % modu

	var tmpr2, tmpr3 int
	if p == 2 && q >= 3 {
		tmpr2 = 1
	} else {
		tmpr2 = -1
	}
	sign := tmpr2
	if sign == 1 || eq1%2 == 0 {
		tmpr3 = 1
	} else {
		tmpr3 = -1
	}
	pm := tmpr3
	return int(math.Pow(float64(p), float64(e0))) * pm * choose_prod % modu
}

func get_crt_root(a_list, n_list []int) int {
	n_mul := reduce(Mul, n_list)
	root_list := Map2(a_list, n_list, func(a, n int) int { return a * n_mul / n * (n_mul / n % a) })
	var res_tmp int
	for i := 0; i < len(root_list); i++ {
		res_tmp += root_list[i]
	}
	return res_tmp % n_mul
}

func in(a int, list [][]int) bool {
	var cnt int = 0
	for i := 0; i < len(list); i++ {
		for _, b := range list[i] {
			if b == a {
				cnt++
			}
		}
	}
	if cnt != 0 {
		return true
	}
	return false
}

func Map(vs []int, f func(int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func Map2(vs, vs2 []int, f func(int, int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v, vs2[i])
	}
	return vsm
}

func Map3(a, b, c []int, f func(int, int, int) int) []int {
	vsm := make([]int, len(a))
	for i := range a {
		vsm[i] = f(a[i], b[i], c[i])
	}
	return vsm
}

func reduce(f func(int, int) int, vs []int) int {
	r := vs[0]
	for _, v := range vs[1:] {
		r = f(r, v)
	}
	return r
}

func Mul(a, b int) int {
	return a * b
}
