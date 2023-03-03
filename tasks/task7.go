package tasks

import (
	"fmt"
	"strconv"
	"sync"
)

func Task7() {
	var wg sync.WaitGroup
	f := func(number int, mp map[int]int, mx *sync.Mutex) {
		mx.Lock()
		mp[number] = number * number
		wg.Done()
		mx.Unlock()
	}

	var syn *sync.Mutex = &sync.Mutex{}
	var n int = 5
	wg.Add(n)
	var m map[int]int = make(map[int]int)
	for x := 0; x < n; x++ {
		go f(x, m, syn)
	}
	wg.Wait()
	for key, value := range m {
		fmt.Println("Key:" + strconv.Itoa(key) + " value:" + strconv.Itoa(value))
	}

	g := func(number int, mp map[int]int, mx sync.RWMutex) {
		mx.RLock()
		fmt.Printf("value: %v", mp[number])
		mx.RUnlock()

		mx.Lock()
		mp[number+20] = number
		wg.Done()
		mx.Unlock()

	}
	wg = sync.WaitGroup{}
	wg.Add(n)
	var rwmx sync.RWMutex

	for x := 0; x < n; x++ {
		go g(x, m, rwmx)
	}
	wg.Wait()

	for key, value := range m {
		fmt.Println("Key:" + strconv.Itoa(key) + " value:" + strconv.Itoa(value))
	}
}
