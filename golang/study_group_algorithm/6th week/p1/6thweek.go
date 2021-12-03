package solution

import "fmt"

func solution(a []int, b []int) int {
	var n int
	fmt.Scanln(&n)
	temp := make([]int, n)
	var answer int
	for i := 0; i < len(temp); i++ {
		temp[i] = a[i] * b[i]
	}
	for j := 0; j < len(temp); j++ {
		answer += temp[j]
	}
	return answer
}
