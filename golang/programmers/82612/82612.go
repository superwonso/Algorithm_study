package main

func solution(price int, money int, count int) int {
	var sum int
	nprice := make([]int, count)
	for i := 0; i < count; i++ {
		nprice[i] = (i + 1) * price
	}
	for i := 0; i < len(nprice); i++ {
		sum += nprice[i]
	}
	if money-sum < 0 {
		return sum - money
	} else {
		return 0
	}
}
