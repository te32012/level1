package tasks

import "fmt"

func Task12() {
	var array []string = []string{"cat", "cat", "dog", "cat", "tree"}
	var mp map[string]bool = make(map[string]bool)
	for _, v := range array {
		mp[v] = true
	}
	fmt.Println(mp)
}
