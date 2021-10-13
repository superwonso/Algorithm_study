package main //proof to bf

import (
	"fmt"
)

func main() {
	const n int = 4
	const m int = 3
	var result = 0
	/* fmt.Scanln(&n, &m)

	cases := make([][]int, n+1)
	for aa := range cases {
		cases[aa] = make([]int, m+1)
	} // to declare 2nd dimension array

	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &cases[i][j])
		}
	}
	*/ //error occur

	var cases = [20][20]int{
		{3, 4, 1, 2},
		{4, 3, 2, 1},
		{3, 1, 4, 2},
	}
	var x, y, z, w int
	var flag int
	for x = 1; x <= n; x++ { // to count possibility
		for y = 1; y <= n; y++ { // to count possibility
			flag = 0
			for w = 0; w < m; w++ { // to proof win in all tests
				result_x := 0
				result_y := 0
				for z = 0; z < n; z++ { // to probe ranking
					if cases[w][z] == x {
						result_x = z
					}
					if cases[w][z] == y {
						result_y = z
					}
				}
				if result_x < result_y {
					flag++
				}
			}
			if flag == m {
				result++
			}
		}
	}
	fmt.Printf("%d", result)
}
