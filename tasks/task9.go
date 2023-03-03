package tasks

import "fmt"

const N int = 10

func parse(array [N]int) chan int {
	var output chan int = make(chan int)
	go func() {
		for _, value := range array {
			output <- value
		}
		close(output)
	}()
	return output
}

func sqr(input chan int) chan int {
	var output chan int = make(chan int)
	go func() {
		for value := range input {
			output <- value * value
		}
		close(output)
	}()
	return output
}

func Task9() {
	var array [N]int = [N]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var first chan int
	var second chan int
	first = parse(array)
	second = sqr(first)
	for v := range second {
		fmt.Println(v)
	}
}
