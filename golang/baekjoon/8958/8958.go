package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(r, &tc)

	score := make([]int, tc)
	str := make([]string, tc)

	for i := 0; i < tc; i++ {
		var tmp, continuity int
		fmt.Fscan(r, &str[i])
		for j := 0; j < len(str[i]); j++ {
			if fmt.Sprintf("%c", str[i][j]) == "O" {
				continuity++
				tmp += continuity
			} else {
				continuity = 0
			}
		}
		score[i] = tmp
	}
	for i := 0; i < tc; i++ {
		fmt.Println(score[i])
	}
}
