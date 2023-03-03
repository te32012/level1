package tasks

import "fmt"

func Task13() {
	var x int = 1
	var y int = -1
	fmt.Printf("x = %v, y = %v \n", x, y)

	x, y = y, x
	fmt.Printf("x = %v, y = %v \n", x, y)
}
