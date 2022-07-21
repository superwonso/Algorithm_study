package main

func sum(a []int) int {
	var __sum int
	for _, v := range a {
		__sum += v
	}
	return __sum
}
