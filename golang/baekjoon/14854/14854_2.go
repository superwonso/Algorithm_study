package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var r = bufio.NewReader(os.Stdin)

func main() {
	var t int
	var result []int
	var D int = 142857
	fmt.Fscan(r, &t)
	for i := 0; i < t; i++ {
		var n, r1 int
		fmt.Fscan(r, &n, &r1)
		var a int = comb_pe(n, r1, 3, 3)
		fmt.Println("a")
		var b int = comb_pe(n, r1, 11, 1)
		fmt.Println("b")
		var c int = comb_pe(n, r1, 13, 1)
		fmt.Println("c")
		var d int = comb_pe(n, r1, 37, 1)
		fmt.Println("d")
		fmt.Println(a, b, c, d)
		tmpr := ((((a * 137566) % D) + ((b * 103896) % D) + ((c * 109890) % D) + (d * 77220 % D)) % D)
		fmt.Println(" ")
		fmt.Println(tmpr)
		fmt.Println(" ")
		//fmt.Println(a)
		result = append(result, tmpr%D)
	}
	for i := 0; i < t; i++ {
		fmt.Println(result[i])
	}
}

func precom(p, e int) []int {
	pe := int(math.Pow(float64(p), float64(e)))
	var val []int = []int{1}
	for idx := 1; idx < pe; idx++ {
		if idx%p == 1 {
			val = append(val, val[len(val)-1]*idx)
		} else {
			val = append(val, val[len(val)-1])
		}
	}
	return val
}

func mod_pow(c, n, d int) int {
	var res int = 1
	for n > 0 {
		if n&1 == 1 {
			res *= c % d
		}
		c *= c % d
		n >>= 1
	}
	return res
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return quotient, remainder
}

func comb(n, k, p int) int {
	if k == 0 {
		return 1
	}
	if n == 0 {
		return 0
	}
	r, s, t := 1, 1, 1
	for i := 2; i <= k; i++ {
		r *= i % p
		if i == k {
			s = mod_pow(r, p-2, p)
		}
		if i == n-k {
			t = mod_pow(r, p-2, p)
		}
	}
	return r * s % p * t % p
}

func lucas(n, k, p int) int {
	var ni []int
	var ki []int
	for n > 0 {
		ni = append(ni, n%p)
		n /= p
	}
	for k > 0 {
		ki = append(ki, k%p)
		k /= p
	}
	for len(ni) > len(ki) {
		ki = append(ki, 0)
	}
	ret := 1
	for i := 0; i < len(ni); i++ {
		ret = (ret * comb(ni[i], ki[i], p)) % p
	}
	return ret
}

func comb_pe(n, k, p, e int) int {
	if e == 1 {
		return lucas(n, k, p)
	}
	nt, nm := fact_n_pe(n, p, e)
	kt, km := fact_n_pe(k, p, e)
	nkt, nkm := fact_n_pe(n-k, p, e)
	fmt.Println("nt nm kt km nkt nkm", nt, nm, kt, km, nkt, nkm)
	pe := int(math.Pow(float64(p), float64(e)))
	fmt.Println("p^e", pe)
	t := nt - kt - nkt
	fmt.Println("t", t)
	if t >= e {
		return 0
	}
	var i int
	for i = 0; i < pe; i++ {
		if (nkm*km*i)%pe == 1 {
			break
		}
	}
	fmt.Println("(int(math.Pow(float64(p), float64(t))) * nm * i) % pe", int(math.Pow(float64(p), float64(t)))*nm*i%pe)
	return int(math.Pow(float64(p), float64(t))) * nm * i % pe
}

func doo(n, p, pe int, val []int) (int, int) {
	if n < p {
		return 0, val[n]
	}
	t := n / p
	nt, nm := doo(t, p, pe, val)
	tp := t + nt
	kp, rp := divmod(n, pe)
	m := ((nm * mod_pow(val[pe-1], kp, pe) % pe) * val[rp]) % pe
	return tp, m
}

func fact_n_pe(n, p, e int) (int, int) {
	val := precom(p, e)
	pe := int(math.Pow(float64(p), float64(e)))
	return doo(n, p, pe, val)
}

/*
	do := func(n int) (int, int) {
		if n < p {
			return 0, val[n]
		}
		t := n / p
		nt, nm := do(t) //doo(t, p, pe, val)
		tp := t + nt
		kp, rp := divmod(n, pe)
		m := (((nm * mod_pow(val[pe-1], kp, pe)) % pe) * val[rp]) % pe
		return tp, m
	}
*/
