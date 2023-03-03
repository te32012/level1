package tasks

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func getterData(n int, channal chan int) {
	for {
		c := <-channal
		fmt.Printf("Потребитель %v время %v \n", n, c)
	}
}

func Task4() {
	var n int
	fmt.Println("введите число потоков")
	fmt.Scanf("%v \n", &n)

	channal := make(chan int)
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT)
	var end bool = false
	go func() {
		sig := <-s
		fmt.Println("\n программа заверешена ")
		fmt.Println(sig)
		end = true
	}()
	for x := 0; x < n; x++ {
		go getterData(x, channal)
	}

	for {
		channal <- time.Now().Second()
		time.Sleep(time.Second * 2)
		if end {
			break
		}
	}
}
