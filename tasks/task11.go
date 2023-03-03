package tasks

import (
	"fmt"
)

type Set struct {
	data map[int]bool
}

func (s *Set) Set(array []int) {
	s.data = make(map[int]bool)
	for _, v := range array {
		s.data[v] = true
	}
}

func (s *Set) Add(d int) {
	if len(s.data) == 0 {
		s.data = make(map[int]bool)
	}
	s.data[d] = true
}

func (s *Set) contains(s1 Set) Set {
	var ans Set
	for k := range s.data {
		if s1.data[k] {
			ans.Add(k)
		}
	}
	return ans
}

func (s *Set) print() {
	for k, _ := range s.data {
		fmt.Printf(" %v ", k)
	}
}

func Task11() {
	var s1 Set
	var s2 Set
	s1.Set([]int{2, 4, 5, 7, 8, 16})
	s2.Set([]int{5, 7, 9})
	var s3 = s1.contains(s2)
	s3.print()
	fmt.Println()
}
