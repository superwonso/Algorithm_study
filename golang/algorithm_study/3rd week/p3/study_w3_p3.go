package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int = 10
	var k int = 3
	// fmt.Scanln(&n, &k)
	// temp, _ := intScanln(n)
	temp_sum := 0
	temp := [10]int{12, 15, 11, 20, 25, 10, 20, 19, 13, 15}
	result := make([]int, n+k+1)
	for i := 0; i < n-k; i++ {
		for j := i; j < i+k; j++ {
			temp_sum = temp_sum + temp[j]
		}
		result[i] = temp_sum // number of cases of result is n+k+1
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	fmt.Println(result[n-k])
}

func intScanln(n int) ([]int, error) {
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err := fmt.Scanln(y...)
	x = x[:n]
	return x, err
}
