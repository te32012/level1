package tasks

import "fmt"

func remove(slice []int, index int) []int {
	var ans []int = []int{}
	if len(slice) > index && index >= 0 {
		ans = append(ans, slice[:index]...)
		ans = append(ans, slice[index+1:]...)
	}
	return ans
}

func Task23() {
	var slice []int = []int{1, 2, 4, 5, 8, 16, 32, 33, 64, 127, 128}
	fmt.Println(remove(slice, 0))
	fmt.Println(remove(slice, 5))
	fmt.Println(remove(slice, len(slice)-1))

}
