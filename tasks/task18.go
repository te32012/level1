package tasks

import (
	"fmt"
	"sync"
)

type Counter struct {
	Count int
	Mx    sync.RWMutex
}

func (c *Counter) add(n int) {
	c.Mx.Lock()
	c.Count += n
	c.Mx.Unlock()
}
func (c *Counter) Counter() {
	c.Count = 0
	c.Mx = sync.RWMutex{}
}
func (c *Counter) read() int {
	c.Mx.RLock()
	tmp := c.Count
	c.Mx.RUnlock()
	return tmp
}

func Task18() {
	var c Counter
	var wg sync.WaitGroup
	wg.Add(10)
	c.Counter()
	for x := 0; x < 10; x++ {
		go func(y int) {
			c.add(y)
			tmp := c.read()
			fmt.Printf("id=%v, read=%v \n", y, tmp)
			wg.Done()
		}(x)
	}
	wg.Wait()
	fmt.Println(c.read())
}
