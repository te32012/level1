package tasks

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var sigma int = 0
var mutex sync.Mutex
var waitGroup sync.WaitGroup

func sumsqrt(slice []int, first int, last int) {
	if last-first > 1 {
		if (last-first)%2 == 0 {
			if (last - first) == 2 {
				// fmt.Println("gorutine one - first : " + strconv.Itoa(first) + "last:" + strconv.Itoa(first+1))
				// fmt.Println("gorutine two - first : " + strconv.Itoa(first+1) + "last:" + strconv.Itoa(last))
				go sumsqrt(slice, first, first+1)
				go sumsqrt(slice, first+1, last)
			} else {
				// fmt.Println("gorutine one - first : " + strconv.Itoa(first) + "last:" + strconv.Itoa(last/2))
				// fmt.Println("gorutine two - first : " + strconv.Itoa((last / 2)) + "last:" + strconv.Itoa(last))
				go sumsqrt(slice, first, (last / 2))
				go sumsqrt(slice, (last / 2), last)
			}
		} else {
			// fmt.Println("gorutine one - first : " + strconv.Itoa(first) + "last:" + strconv.Itoa(last-1))
			// fmt.Println("gorutine two - first : " + strconv.Itoa(last-1) + "last:" + strconv.Itoa(last))
			go sumsqrt(slice, first, last-1)
			go sumsqrt(slice, last-1, last)
		}
	} else {
		var n int = slice[first]
		mutex.Lock()
		sigma = sigma + n*n
		waitGroup.Done()
		mutex.Unlock()
	}
}

func Task3() {
	var slice []int = []int{2, 4, 6, 8, 10}
	go sumsqrt(slice, 0, len(slice))
	s := make(chan os.Signal, 1)
	waitGroup.Add(len(slice))
	waitGroup.Wait()
	fmt.Println(sigma)
	signal.Notify(s, syscall.SIGINT)
	<-s
}
