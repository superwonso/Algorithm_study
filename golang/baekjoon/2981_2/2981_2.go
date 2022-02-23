package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var N int
	fmt.Scanln(&N)
	m := make([]int, N)
	result := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &m[i])
	}
	sort.Sort(sort.IntSlice(m))
	tmp := m[1] - m[0]

	for i := 2; i < N; i++ {
		tmp = gcd(tmp, m[i]-m[i-1])
	}

	for i := 2; tmp >= i*i; i++ {
		if tmp%i == 0 {
			result = push_back(result, i)
			result = push_back(result, tmp/i)
		}
	}
	result = push_back(result, tmp)
	sort.Sort(sort.IntSlice(result))
	result = unique(result)
	if result[0] == 0 {
		for i := 1; i < len(result); i++ {
			fmt.Printf("%d ", result[i])
		}
	} else {
		for i := 0; i < len(result); i++ {
			fmt.Printf("%d ", result[i])
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

//1713
func push_back(arr []int, val int) []int {
	arr = append(arr, val)
	return arr
}

//https://www.golangprograms.com/remove-duplicate-values-from-slice.html
func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
