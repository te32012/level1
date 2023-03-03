package tasks

import "fmt"

type all interface{}

func Task14() {

	f := func(p all) {
		fmt.Printf("type of parametr - %T \n", p)
	}
	f("a")
	f(-1)
	f(true)
	var ch chan int = make(chan int)
	f(ch)
}
