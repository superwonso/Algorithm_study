package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var N, res int
	var tmp string
	fmt.Fscanln(r, &N)
	fmt.Fscanf(r, "%s", &tmp)
	tmps := strings.Split(tmp, "")
	for i := 0; i < N; i++ {
		tp, _ := strconv.Atoi(tmps[i])
		res += tp
	}
	fmt.Println(res)
}
