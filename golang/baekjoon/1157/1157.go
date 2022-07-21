package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	res := make([]int, 26)
	var input string
	var trash int
	fmt.Fscanln(r, &input)
	for x := 0; x < len(input); x++ {
		i := input[x]
		if (i == 65) || (i == 97) {
			res[0] += 1
		} else if (i == 66) || (i == 98) {
			res[1] += 1
		} else if (i == 67) || (i == 99) {
			res[2] += 1
		} else if (i == 68) || (i == 100) {
			res[3] += 1
		} else if (i == 69) || (i == 101) {
			res[4] += 1
		} else if (i == 70) || (i == 102) {
			res[5] += 1
		} else if (i == 71) || (i == 103) {
			res[6] += 1
		} else if (i == 72) || (i == 104) {
			res[7] += 1
		} else if (i == 73) || (i == 105) {
			res[8] += 1
		} else if (i == 74) || (i == 106) {
			res[9] += 1
		} else if (i == 75) || (i == 107) {
			res[10] += 1
		} else if (i == 76) || (i == 108) {
			res[11] += 1
		} else if (i == 77) || (i == 109) {
			res[12] += 1
		} else if (i == 78) || (i == 110) {
			res[13] += 1
		} else if (i == 79) || (i == 111) {
			res[14] += 1
		} else if (i == 80) || (i == 112) {
			res[15] += 1
		} else if (i == 81) || (i == 113) {
			res[16] += 1
		} else if (i == 82) || (i == 114) {
			res[17] += 1
		} else if (i == 83) || (i == 115) {
			res[18] += 1
		} else if (i == 84) || (i == 116) {
			res[19] += 1
		} else if (i == 85) || (i == 117) {
			res[20] += 1
		} else if (i == 86) || (i == 118) {
			res[21] += 1
		} else if (i == 87) || (i == 119) {
			res[22] += 1
		} else if (i == 88) || (i == 120) {
			res[23] += 1
		} else if (i == 89) || (i == 121) {
			res[24] += 1
		} else if (i == 90) || (i == 122) {
			res[25] += 1
		} else {
			trash++
		}
	}
	var tmp, x, record int = 0, 0, 0
	for i := 0; i < 26; i++ {
		if res[i] > tmp {
			tmp = res[i]
			record = i
		}
	}
	for i := 0; i < 26; i++ {
		if res[i] == tmp {
			x++
		}
	}
	if x > 1 {
		fmt.Println("?")
	} else {
		fmt.Printf("%c", record+65)
	}

}
