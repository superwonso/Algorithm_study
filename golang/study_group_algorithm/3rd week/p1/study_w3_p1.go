package main

import (
	"fmt"
	"reflect"
	"sort"
)

// first, ascending sort then find intersection
func main() {
	var temp_a, temp_b int
	var b []int

	fmt.Scanln(&temp_a)
	a, _ := intScanln(temp_a)

	fmt.Scanln(&temp_b)
	b, err := intScanln(temp_b)

	if err != nil {
		fmt.Println(err)
		return
	}

	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.IntSlice(b))
	fmt.Println(Intersect(a, b)) // complexitsy : O(n^2)
}

func Intersect(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	av := reflect.ValueOf(a)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		if contains(b, el) {
			set = append(set, el)
		}
	}

	return set
}

func contains(a interface{}, e interface{}) bool {
	v := reflect.ValueOf(a)

	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == e {
			return true
		}
	}
	return false
}

func intScanln(n int) ([]int, error) {
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err := fmt.Scanln(y...)
	x = x[:n]
	return x, err
}
