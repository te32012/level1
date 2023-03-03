package tasks

import (
	"fmt"
	"sync"
	"time"
)

func Task5() {
	var n int
	fmt.Println("Введите число секунд -- время работы программы")
	fmt.Scanln(&n)
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	var timer time.Timer
	timer = *time.NewTimer(time.Duration(n) * time.Second)
	var channal chan time.Time = make(chan time.Time)

	go func() {
		for {
			time, isclose := <-channal
			if isclose {
				fmt.Printf("Прогресс выполнения программы - %v текущее время \n", time)
			} else {
				fmt.Println("Программа завершила свою работу")
				waitGroup.Done()
				return
			}
		}
	}()

outerloop:
	for {
		select {
		case <-timer.C:
			fmt.Println("Время выполнения программы заканчивается")
			close(channal)
			break outerloop
		default:
			channal <- time.Now()
		}
	}
	waitGroup.Wait()
}
