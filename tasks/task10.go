package tasks

import (
	"fmt"
	// "math"
)

func Task10() {
	var temp []float64 = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	var mp map[int][]float64 = make(map[int][]float64)
	var k int

	for _, tmp := range temp {
		k = (int(tmp) / 10) * 10
		mp[k] = append(mp[k], tmp)
	}
	fmt.Println(mp)
}
