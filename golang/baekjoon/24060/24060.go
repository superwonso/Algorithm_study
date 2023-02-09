package main

import "fmt"

var a = make([]int, 500001)
var tmp = make([]int, 500001)
var N, cnt, k int = 0, 0, 0
var result int = -1

func main() {
	fmt.Scanf("%d %d", &N, &k)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
	}
	mergeSort(a, 0, N-1)
	fmt.Println(result)
}

func mergeSort(a []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(a, left, mid)
		mergeSort(a, mid+1, right)
		merge(a, left, mid, right)
	}
}

func merge(a []int, left, mid, right int) {
	i := left
	j := mid + 1
	t := 1
	for i <= mid && j <= right {
		if a[i] < a[j] {
			tmp[t] = a[i]
			t++
			i++
		} else {
			tmp[t] = a[j]
			t++
			j++
		}
	}
	for i <= mid {
		tmp[t] = a[i]
		t++
		i++
	}
	for j <= right {
		tmp[t] = a[j]
		t++
		j++
	}
	i = right
	t = 1
	for i <= right {
		a[i] = tmp[t]
		i++
		t++
		cnt++
		if cnt == k {
			result = a[i-1]
			break
		}
	}
}
