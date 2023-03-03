package tasks

import "fmt"

// Про наследование в go можно почитать например тут
// https://github.com/MaksimDzhangirov/oop-go/blob/master/docs/inheritance.md

type Human struct {
	firstname string
	lastname  string
}

type Action struct {
	Human
	age int
}

func (act *Action) print() {
	fmt.Println(act.firstname)
	fmt.Println(act.lastname)
	fmt.Println(act.age)
}

func Task1() {
	var h Human = Human{firstname: "anna", lastname: "anna"}
	var c Action = Action{Human: h, age: 15}
	c.print()
}
