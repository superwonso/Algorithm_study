package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N int
	var result_MAX, result_MIN float64
	fmt.Scan(&N)
	slice := make([]float64, N)
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &slice[i])
	}
	sort.Sort(sort.Float64Slice(slice))
	result_MAX = slice[len(slice)-1]
	result_MIN = slice[0]
	fmt.Println(int(result_MIN), int(result_MAX))
}
