package main

import (
	"fmt"
)

func main() {
	var dice1, dice2, dice3 int
	var N int
	fmt.Scanln(&N)
	Price := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&dice1, &dice2, &dice3)
		if dice1 == dice2 && dice2 == dice3 {
			Price[i] = 10000 + ((dice1) * 1000)
		} else if dice1 == dice2 && dice1 != dice3 {
			Price[i] = 1000 + ((dice1) * 100)
		} else if dice1 == dice3 && dice1 != dice2 {
			Price[i] = 1000 + ((dice1) * 100)
		} else if dice2 == dice3 && dice1 != dice2 {
			Price[i] = 1000 + ((dice2) * 100)
		} else {
			var temp_dice int = 0
			if dice1 > dice2 {
				if dice1 > dice3 {
					temp_dice = dice1
				} else if dice3 > dice1 {
					temp_dice = dice3
				}
			} else if dice2 > dice1 {
				if dice2 > dice3 {
					temp_dice = dice2
				} else if dice3 > dice2 {
					temp_dice = dice3
				}
			}
			Price[i] = temp_dice * 100
		}
	}
	for i := 0; i < len(Price); i++ {
		for j := 0; j < i; j++ {
			if Price[i] < Price[j] {
				Price[i], Price[j] = Price[j], Price[i]
			}
		}
	}
	fmt.Println(Price[len(Price)-1])
}
