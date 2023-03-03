package tasks

import (
	"fmt"
)

func Task17() {
	var array []int = []int{1, 2, 4, 5, 8, 15, 16, 32, 64, 128, 128}
	fmt.Println(binarysearch(array, 5, 0, len(array)))
}

func binarysearch(array []int, element int, leftindex int, rightindex int) int {
	middle := leftindex + ((rightindex - leftindex) / 2)
	//fmt.Print(middle)
	switch {
	case (rightindex-leftindex)/2 > 0 && (element == array[middle]):
		//fmt.Println("find")
		return middle
	case (rightindex-leftindex)/2 > 0 && (element > array[middle]):
		//fmt.Println("right")
		return binarysearch(array, element, middle, rightindex)
	case (rightindex-leftindex)/2 > 0 && (element < array[middle]):
		//fmt.Println("left")
		return binarysearch(array, element, leftindex, middle)
	default:
		//fmt.Println("notfind")
		return -1
	}
}
