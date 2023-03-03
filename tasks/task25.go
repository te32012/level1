package tasks

import (
	"fmt"
	"time"
)

func Task25() {
	var timer time.Timer = *time.NewTimer(5 * time.Second)
	<-timer.C
	fmt.Println("stop after 5 second")
}
