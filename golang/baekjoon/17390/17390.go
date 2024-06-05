package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	params := strings.Fields(line)
	n, _ := strconv.Atoi(params[0])
	q, _ := strconv.Atoi(params[1])

	line, _ = reader.ReadString('\n')
	a := make([]int, n)
	elements := strings.Fields(line)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(elements[i])
	}

	sort.Ints(a)

	d := make([]int, n+1)
	for i := 0; i < n; i++ {
		d[i+1] = d[i] + a[i]
	}

	for i := 0; i < q; i++ {
		line, _ = reader.ReadString('\n')
		query := strings.Fields(line)
		l, _ := strconv.Atoi(query[0])
		r, _ := strconv.Atoi(query[1])
		fmt.Fprintln(writer, d[r]-d[l-1])
	}
}
