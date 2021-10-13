package main

import (
	"fmt"
)

func main() {

	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	var chair, attacked []int
	for i := 1; i <= n; i++ {
		chair = append(chair, i)
	}

	var temp int
	for len(chair) > 0 {
		temp = (temp + k - 1) % len(chair)
		attacked = append(attacked, chair[temp])
		chair = append(chair[:temp], chair[temp+1:]...)
	}

	for i, v := range attacked {
		if i == 0 {
			fmt.Printf("<%d", v)
		} else {
			fmt.Printf(", %d", v)
		}
	}
	fmt.Println(">")
}
