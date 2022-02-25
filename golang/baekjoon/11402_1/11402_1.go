package main

import (
	"fmt"
)

func main() {
	var n, k, m int
	fmt.Scanln(&n, &k, &m)
	var result int = 1
	for n > 0 || k > 0 {
		result *= Combination(n%m, k%m) % m
		n /= m
		k /= m
	}
	/*converted_n := make([]int, 100)
	converted_k := make([]int, 100)
	converted_n = change_digits(n, m)
	converted_k = change_digits(k, m)
	result := Combination(converted_n[0], converted_k[0])
	for i := 1; i < len(converted_k); i++ {
		result = result * Combination(converted_n[i], converted_k[i])
		if converted_k[len(converted_k)-1] == 0 {
			break
		} else if converted_n[i] == 0 && converted_k[i] == 0 {
			break
		}
	}
	*/
	//if m == 0 {
	//fmt.Println(result)
	//} else {
	fmt.Println(result % m)
	//}
}

func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return (n * factorial(n-1))
}

func Permutation(n, k int) int {
	if n == 0 || k == 0 {
		return 1
	}
	return factorial(n) / factorial(n-k)
}

func Combination(n, k int) int {
	if n < k {
		return 0
	} else if n == 0 || k == 0 {
		return 1
	}
	return Permutation(n, k) / factorial(k)
}

/*func change_digits(n, m int) []int {
	result := make([]int, 100)
	if n == 0 || m == 0 {
		return result
	}
	var i int = 0
	for n > 0 {
		result[i] = n % m
		n = n / m
		i++
	}
	//sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result
}
*/
