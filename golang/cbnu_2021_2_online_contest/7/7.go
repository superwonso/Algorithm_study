package main

import "fmt"

func main() {
	var tc, t int
	result := [...]string{"NA", "MMYY", "YYMM", "AMBIGOUS"}
	fmt.Scanln(&t)
	for tc = 1; tc <= t; tc++ {
		var c [4]rune
		result_num := 0
		fmt.Scanf("%s", c)
		frontNum := (c[0]-'0')*10 + (c[1] - '0')
		backNum := (c[2]-'0')*10 + (c[3] - '0')
		if frontNum < 13 && backNum > 0 {
			result_num += 1
		}
		if backNum < 13 && frontNum > 0 {
			result_num += 2
		}
		fmt.Printf("#%d %s\n", tc, result[result_num])
	}
}
