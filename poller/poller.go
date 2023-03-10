package poller

import (
	"log"
	"time"
)

func Run(interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			go getPendingTasks()
			go deployRedyTasks()
			go timeoutTasks()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func getPendingTasks() {
	log.Println("getting pending tasks")
}
func deployRedyTasks() {
	log.Println("deploying ready executions")
}
func timeoutTasks() {
	log.Println("killing timed out tasks")
}
