package tasks

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Первый вариант - через переменную типа bool
Второй вариант - через канал <-
Третий вариант - через специальную переменную канала
Четвертый вариант - по таймеру
Пятый вариант - через контест
*/
func Task6() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)
	var isclose bool = false
	go func(isclose *bool) {
		for {
			if *isclose {
				fmt.Println("Закрытие горутины номер 1")
				waitGroup.Done()
				return
			}
		}
	}(&isclose)
	isclose = true

	var channal chan int = make(chan int)
	go func() {
		for {
			select {
			case <-channal:
				fmt.Println("Закрытие горутины номер 2")
				waitGroup.Done()
				return
			}
		}
	}()
	channal <- 1
	var channal2 chan int = make(chan int)
	go func() {
		for {
			_, ok := <-channal2
			if !ok {
				fmt.Println("Закрытие горутины номер 3")
				waitGroup.Done()
				return
			}
		}
	}()
	channal2 <- 1
	close(channal)
	close(channal2)

	var timer time.Timer = *time.NewTimer(time.Second * 3)

	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("Закрытие горутины номер 4")
				waitGroup.Done()
				return
			}
		}
	}()
	ctx, end := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Закрытие горутины номер 5")
				waitGroup.Done()
				return
			}
		}
	}()
	end()
	waitGroup.Wait()
}
