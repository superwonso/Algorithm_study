package main

import (
	"fmt"
)

func main() {
	var N int
	var max_val float64 = 0.0
	fmt.Scan(&N)
	value := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&value[i])
		if value[i] > max_val {
			max_val = value[i]
		}
	}
	var sum float64
	for i := 0; i < N; i++ {
		value[i] = value[i] / max_val * 100
		sum += value[i]
	}
	new_avg := sum / float64(len(value))
	fmt.Println(float64(new_avg))

}
