package scheduler

import (
	"fmt"
	"time"
)

type Scheduler struct {
	jobs []string
}

func New() Scheduler {
	return Scheduler{}
}

func (s Scheduler) Start(done <-chan bool) {
	fmt.Println("Scheduler start")

	for {
		select {
		case d := <-done:
			fmt.Println("exiting...", d)
			return
		default:
			now := time.Now()
			fmt.Println("Scheduler now: ", now)
			time.Sleep(1 * time.Second)
		}
	}
}
