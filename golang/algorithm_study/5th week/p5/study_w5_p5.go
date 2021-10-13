package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	temp := bufio.NewReader(os.Stdin)
	s, err := temp.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	output := strings.Replace(s, "zero", "0", -1)
	output = strings.Replace(output, "one", "1", -1)
	output = strings.Replace(output, "two", "2", -1)
	output = strings.Replace(output, "three", "3", -1)
	output = strings.Replace(output, "four", "4", -1)
	output = strings.Replace(output, "five", "5", -1)
	output = strings.Replace(output, "six", "6", -1)
	output = strings.Replace(output, "seven", "7", -1)
	output = strings.Replace(output, "eight", "8", -1)
	output = strings.Replace(output, "nine", "9", -1)
	fmt.Println(output)
}
