package main

import (
	"fmt"
)

func main() {
	array := [5][5]int{
		{0, 0, 0, 0, 0}, {0, 0, 1, 0, 3}, {0, 2, 5, 0, 1}, {4, 2, 4, 4, 2}, {3, 5, 1, 3, 1},
	}
	// temp := bufio.NewReader(os.Stdin)
	const N int = 5
	/* fmt.Fscan(temp, &N)
	array := make([][]int, N)
	 for a := range array {
		array[a] = make([]int, N)
	}
	for j := 0; j < N; j++ {
		for i := 0; i < N; i++ {
			fmt.Scanf("%d", &array[i][j])
		}
	} // empty slice(=array) is filled to "0"*/
	move_list := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &move_list[i])
	}
	basket := make([]int, N)
	move(move_list, array, basket, N)
	answer := 0
	clear(basket, N, answer)
	fmt.Println(2 * answer)
}

func move(move []int, from [5][5]int, target []int, N int) {
	for i := 0; i < N; i++ {
		if move[i] == 0 {
			i = i + 1
			continue
		}
		for j := 0; j < N; j++ {
			if from[move[i]][j] == 0 {
				j = j + 1
				continue
			}
			target[move[i]] = from[move[i]][j]
			from[move[i]][j] = 0
		}
	}

}

func clear(basket []int, N int, answer int) {
	for i := 1; i < N; i++ {
		if basket[i-1] == basket[i] {
			for j := i; j < N; j++ {
				basket[j-1] = basket[j+1]
				basket[j] = basket[j+2]
				basket[j+1] = 0
				basket[j+2] = 0
				answer++
			}

		}
	}
}
