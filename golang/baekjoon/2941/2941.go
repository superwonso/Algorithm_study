package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var r = bufio.NewReader(os.Stdin)

func main() {
	var s string
	var res []string
	fmt.Fscan(r, &s)
	s = strings.Replace(s, "c=", "!", -1)
	s = strings.Replace(s, "c-", "@", -1)
	s = strings.Replace(s, "dz=", "#", -1)
	s = strings.Replace(s, "d-", "$", -1)
	s = strings.Replace(s, "lj", "%", -1)
	s = strings.Replace(s, "nj", "^", -1)
	s = strings.Replace(s, "s=", "&", -1)
	s = strings.Replace(s, "z=", "*", -1)
	for i := 0; i < len(s); i++ {
		res = append(res, string(s[i]))
	}
	fmt.Println(len(res))
}
