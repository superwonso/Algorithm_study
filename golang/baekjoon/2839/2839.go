package main

import "fmt"

func main() {
	var sugar int
	var vinyl_5 int = 0
	var vinyl_3 int = 0
	fmt.Scanln(&sugar)
	for true {
		if sugar-5 > 0 {
			sugar -= 5
			vinyl_5++
		} else if sugar-3 > 0 {
			sugar -= 3
			vinyl_3++
		}
		if sugar-5 < 0 && sugar-3 < 0 {
			break
		}
	}
	if sugar != 0 {
		fmt.Println("-1")
	} else {
		fmt.Println(vinyl_3 + vinyl_5)
	}

}
