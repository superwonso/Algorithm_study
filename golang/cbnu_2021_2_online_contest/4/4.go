package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)
	help := make([]string, N)
	var midterm, final, homework float64
	students := make([]int, N)
	score := make([]float64, N)
	var i int
	rank := make([]float64, N)
	for i = 0; i < N; i++ {
		fmt.Scanln(midterm, final, homework)
		score[i] = (0.35 * midterm) + (0.45 * final) + (0.2 * homework)
		students[i] = i + 1
	}
	for j := 0; j < N; j++ {
		for k := 0; k < N; k++ {
			if score[j]-score[k] < 0 {
				i += 1
			}
		}
		rank[i] = score[j]
		i = 0
	}
	temp := (len(students) / 10)
	temp = float64(temp)
	for i := 0; i < N; i++ {
		if rank[i] > 0 && rank[i] <= temp {
			help[i] = "A+"
		}
	}

}
