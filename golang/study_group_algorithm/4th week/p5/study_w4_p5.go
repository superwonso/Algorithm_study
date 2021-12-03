package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	var sum int = 0
	var result int = 0
	temp := bufio.NewReader(os.Stdin)
	fmt.Fscan(temp, &N)
	array := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &array[i])
	}
	for j := 0; j < N; j++ {
		sum += array[j]
	}
	tmp_result := make([]int, N)
	tmp_result[0] = 0
	tmp_result[N-1] = 0
	for k := 1; k < N; k++ {
		if array[k] == 0 {
			tmp_result[k] = array[k-1]
		} else {
			tmp_result[k] = 0
		}
	}
	var tmp_result2 int
	for l := 0; l < N; l++ {
		tmp_result2 += tmp_result[l]
	}
	result = sum - tmp_result2
	fmt.Println(result)

}

//wrong
