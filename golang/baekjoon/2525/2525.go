package main

import (
	"fmt"
)

func main() {
	var h, m int
	var hour, min, add int
	fmt.Scanln(&hour, &min)
	fmt.Scanln(&add)
	h = add / 60
	m = add % 60
	hour = hour + h
	min = min + m
	if min > 59 {
		min = min - 60
		hour++
	}
	if hour > 23 {
		hour = hour - 24
	}
	fmt.Println(hour, min)
}
