package poller

import (
	"time"

	"github.com/drorivry/matter/dao"
	"github.com/drorivry/matter/models"
)

func Run(interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			go deployRedyTasks()
			go timeoutTasks()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func deployRedyTasks() {
	tasks := dao.GetPendingTasks()

	for _, task := range tasks {
		taskEx := models.CreateExecutionFromDefinition(task)
		dao.InsertTaskExecution(taskEx)

		task.ExecutionsCounter++

		if task.ExecutionInterval > 0 {
			task.NextExecutionTime = time.Now().Add(time.Duration(task.ExecutionInterval) * time.Second)
		}

	}
}

func timeoutTasks() {
	//log.Println("killing timed out tasks")
}
