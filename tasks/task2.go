package tasks

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const n = 5

// Классическая реализация стратегии разделяй и властвуй
// с вычислением квадрата в качестве полезной нагрузки
func sqrt(array *[n]int, first int, last int) {
	if last-first > 1 {
		if (last-first)%2 == 0 {
			if (last - first) == 2 {
				fmt.Println("gorutine one - first : " + strconv.Itoa(first) + "last:" + strconv.Itoa(first+1))
				fmt.Println("gorutine two - first : " + strconv.Itoa(first+1) + "last:" + strconv.Itoa(last))
				go sqrt(array, first, first+1)
				go sqrt(array, first+1, last)
			} else {
				fmt.Println("gorutine one - first : " + strconv.Itoa(first) + "last:" + strconv.Itoa(last/2))
				fmt.Println("gorutine two - first : " + strconv.Itoa((last / 2)) + "last:" + strconv.Itoa(last))
				// Вычисляем корни c 0 до last / 2 не включительно
				go sqrt(array, first, (last / 2))
				// Вычисляем корни c last / 2 до last не включительно
				go sqrt(array, (last / 2), last)
			}
		} else {
			fmt.Println("gorutine one - first : " + strconv.Itoa(first) + "last:" + strconv.Itoa(last-1))
			fmt.Println("gorutine two - first : " + strconv.Itoa(last-1) + "last:" + strconv.Itoa(last))
			go sqrt(array, first, last-1)
			go sqrt(array, last-1, last)
		}
	} else {
		var n int = (*array)[first]
		fmt.Println(n * n)
	}
}

func Task2() {
	var array [n]int = [n]int{2, 4, 6, 8, 10}
	go sqrt(&array, 0, len(array))
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT)
	<-s
}
