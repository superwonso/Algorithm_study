package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var N, k int = 0, 0
	fmt.Scanln(&N)
	m := make([]string, N)
	res := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &m[i])
	}
	keys := make(map[int]bool)
	var signal bool = true
	for signal == true {
		for i := 0; i < N; i++ {
			tmp, _ := strconv.Atoi(m[i])
			res[i] = tmp % (set_digits(count_digits(Max(m))))
		}
		for _, value := range res {
			if _, saveValue := keys[value]; !saveValue {
				keys[value] = true
				k++
			} else {
				signal = false
				break
			}
		}
	}
	fmt.Println(k)
}

func set_digits(digits int) int {
	for i := 0; i < digits-1; i++ {
		digits *= 10
	}
	return digits
}

func count_digits(n int) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func Max(arr []string) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	int_max, _ := strconv.Atoi(max)
	return int_max
}
