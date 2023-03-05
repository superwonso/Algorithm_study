package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	r = bufio.NewReader(os.Stdin)
)

func main() {
	var n int
	fmt.Fscanln(r, &n)
	arr := make([]int, n)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	var min int
	var index int
	min = MinIntSlice(arr, min) // minimun value of arr
	for i := 0; i < len(arr); i++ {
		if arr[i] == min {
			index = i
		}
	}
	if arr[index] < 0 {
		arr[index] = 0
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			ans[i] += arr[j]
		}
	}
	var result int
	result = MaxIntSlice(ans, result)
	fmt.Println(result)
}

func MinIntSlice(v []int, tmp int) int {
	x := make([]int, len(v))
	copy(x, v)
	sort.Ints(x)
	tmp = x[0]
	return tmp
}

func MaxIntSlice(v []int, tmp int) int {
	x := make([]int, len(v))
	copy(x, v)
	sort.Ints(x)
	tmp = x[len(x)-1]
	return tmp
}
