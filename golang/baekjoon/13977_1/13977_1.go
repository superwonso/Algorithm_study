//not solved
package main

import "fmt"

var mod int64 = 1000000007
var ans = make([]int64, 4000001)

func main() {
	ans[0] = 1
	ans[1] = 1
	var i int64
	for i = 2; i <= 4000000; i++ {
		ans[i] = ans[i-1] * i % mod
	}
	var t, k int
	var a, b int64
	fmt.Scanln(&t)
	for k = 0; k < t; k++ {
		fmt.Scanln(&a, &b)
		fmt.Println(a, b)
		fmt.Println(ans[a] * mod_func(ans[a-b]*ans[b]%mod, mod-2) % mod)
	}
}

func mod_func(a int64, b int64) int64 {
	if b == 1 {
		return a % mod
	}
	if check_odd(b) == true {
		return mod_func(a, b-1) * a % mod
	} else {
		return mod_func(a*a%mod, b/2)
	}
}

func check_odd(n int64) bool {

	if n%2 != 1 {
		return true
	} else {
		return false
	}

}
