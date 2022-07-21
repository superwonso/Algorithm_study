package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var tc int
	fmt.Fscanln(r, &tc)
	defer w.Flush()

	for i := 0; i < tc; i++ {
		var n int
		fmt.Fscanf(r, "%d ", &n)

		var scores = make([]float64, n)
		var sum, avg float64
		for j := 0; j < n; j++ {
			fmt.Fscanf(r, "%f ", &scores[j])
			sum += scores[j]
		}
		avg = sum / float64(n)

		var proportion float64
		for _, val := range scores {
			if val > avg {
				proportion += (1 / float64(n))
			}
		}
		fmt.Fprintf(w, "%.3f%s\n", math.Round(proportion*100000)/1000, "%")
	}
}
