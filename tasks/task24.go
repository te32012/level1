package tasks

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) Point(x float64, y float64) {
	p.x = x
	p.y = y
}

func distance(p1 Point, p2 Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func Task24() {
	var p1 Point
	var p2 Point
	p1.Point(1, 1)
	p2.Point(2, 2)
	fmt.Println(distance(p1, p2))
}
