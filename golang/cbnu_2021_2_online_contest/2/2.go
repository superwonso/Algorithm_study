package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	temp := bufio.NewReader(os.Stdin)
	s, _ := temp.ReadString('\n')
	output := strings.Replace(s, "a", "A", -1)
	output = strings.Replace(output, "b", "B", -1)
	output = strings.Replace(output, "c", "C", -1)
	output = strings.Replace(output, "d", "D", -1)
	output = strings.Replace(output, "e", "E", -1)
	output = strings.Replace(output, "f", "F", -1)
	output = strings.Replace(output, "g", "G", -1)
	output = strings.Replace(output, "h", "H", -1)
	output = strings.Replace(output, "i", "I", -1)
	output = strings.Replace(output, "j", "J", -1)
	output = strings.Replace(output, "k", "J", -1)
	output = strings.Replace(output, "l", "L", -1)
	output = strings.Replace(output, "n", "N", -1)
	output = strings.Replace(output, "m", "M", -1)
	output = strings.Replace(output, "o", "O", -1)
	output = strings.Replace(output, "p", "P", -1)
	output = strings.Replace(output, "q", "Q", -1)
	output = strings.Replace(output, "r", "R", -1)
	output = strings.Replace(output, "s", "S", -1)
	output = strings.Replace(output, "t", "T", -1)
	output = strings.Replace(output, "u", "U", -1)
	output = strings.Replace(output, "v", "V", -1)
	output = strings.Replace(output, "w", "W", -1)
	output = strings.Replace(output, "x", "X", -1)
	output = strings.Replace(output, "y", "Y", -1)
	output = strings.Replace(output, "z", "Z", -1)
	fmt.Println(output)
}
