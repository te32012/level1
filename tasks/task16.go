package tasks

import "fmt"

func Task16() {
	var array []int = []int{1, 2, 8, 16, 32, 128, 3, 5, 3, 7, 127, 255}
	fmt.Println(quicksort(array))
}
func quicksort(array []int) []int {
	base := array[len(array)/2]
	left := []int{}
	right := []int{}
	for x := 0; x < len(array)/2; x++ {
		switch {
		case array[x] < base:
			left = append(left, array[x])
		case array[x] > base:
			right = append(right, array[x])
		default:
			left = append(left, array[x])
		}
	}
	for x := (len(array) / 2) + 1; x < len(array); x++ {
		switch {
		case array[x] < base:
			left = append(left, array[x])
		case array[x] > base:
			right = append(right, array[x])
		default:
			left = append(left, array[x])
		}
	}
	var sortleft []int
	var sortright []int
	var ans []int
	if len(left) > 0 {
		sortleft = quicksort(left)
	}
	if len(right) > 0 {
		sortright = quicksort(right)
	}
	for _, v := range sortleft {
		ans = append(ans, v)
	}
	ans = append(ans, base)
	for _, v := range sortright {
		ans = append(ans, v)
	}
	return ans
}
