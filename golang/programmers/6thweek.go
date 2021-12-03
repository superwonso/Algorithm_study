func solution(a []int, b []int) int {
	temp := make([]int, n)
	var answer int
	for i := 0; i < len(temp); i++ {
		a[i] * b[i] = temp[i]
	}
	for j := 0; j < len(temp); j++ {
		answer += temp[j]
	}
	return answer
}