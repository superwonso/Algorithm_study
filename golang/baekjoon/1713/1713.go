package main

import (
	"fmt"
	"sort"
)

var recommend [110]int
var duration [110]int

func main() {
	var N, M int = 0, 0
	fmt.Scanln(&N, &M)
	photo := make([]int, N)
	for i := 0; i < M; i++ {
		var x int
		fmt.Scanln(&x)
		recommend[x]++
		for j := 0; j < len(photo); j++ {
			duration[photo[j]]++
		}
		var exist bool = false
		for j := 0; j < len(photo); j++ {
			if photo[j] == x {
				exist = true
				break
			}
		}
		if !exist {
			if len(photo) < N {
				push_back(photo, x)
			} else {
				sort.Slice(photo, compare)
				recommend[photo[len(photo)-1]] = 0
				duration[photo[len(photo)-1]] = 0
				photo[len(photo)-1] = x
			}
		}
		sort.Ints(photo)
		for _, i := range photo {
			fmt.Println(photo[i])
		}
	}
}

func compare(o1, o2 int) bool {
	if recommend[o1] == recommend[o2] {
		if duration[o1] < duration[o2] {
			return true
		} else {
			return false
		}
	} else if recommend[o1] > recommend[o2] {
		return true
	} else {
		return false
	}
}

func push_back(arr []int, val int) []int {
	arr = append(arr, val)
	return arr
}
