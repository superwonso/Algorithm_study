package main

import "fmt"

func main() {
	var T int
	var elect_ra = [3]int{300, 60, 10}
	var cnt_ra = [3]int{0, 0, 0}
	fmt.Scanln(&T)
	for i := 0; i < 3; i++ {
		cnt_ra[i] = T / elect_ra[i]
		T %= elect_ra[i]
	}
	if T != 0 {
		fmt.Println("-1")
	} else {
		for i := 0; i < 3; i++ {
			fmt.Printf("%d ", cnt_ra[i])
		}
	}
}
