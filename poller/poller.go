package poller

import (
	"log"
	"time"

	"github.com/drorivry/matter/dao"
	"github.com/drorivry/matter/initializers"
	"github.com/drorivry/matter/models"

	k8s_client "github.com/drorivry/matter/k8s"
)

func Run(interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			go deployReadyTasks()
			go timeoutTasks()
			go updateTaskStatus()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func deployReadyTasks() {
	tasks := dao.GetPendingTasks()
	for _, task := range tasks {
		log.Println("deploying task ", task.ID)
		DeployJob(task)
	}
}

func DeployJob(task models.TaskDefinition) {
	taskEx := models.CreateExecutionFromDefinition(task)
	jobName := k8s_client.BuildJobName(taskEx)

	taskEx.Status = models.JOB_DEPLOYED

	dao.InsertTaskExecution(taskEx)

	task.ExecutionsCounter++

	// TODO: Move to dao
	if task.ExecutionInterval > 0 {
		task.NextExecutionTime = time.Now().Add(time.Duration(task.ExecutionInterval) * time.Second)
	} else {
		task.Enabled = false
	}

	initializers.GetTaskDefinitionsTable().Save(&task)
	initializers.GetTaskExecutionsTable().Save(&taskEx)
	k8s_client.LaunchK8sJob(&jobName, &taskEx)
}

func timeoutTasks() {
	tasksExecutions := dao.GetTasksToTimeout()

	for _, tasksExecution := range tasksExecutions {
		log.Println("timing out task ", tasksExecution.ID)
		k8s_client.AbortTask(tasksExecution.ID)
		dao.UpdateExecutionStatus(tasksExecution.ID, models.TIMEOUT)
	}
}

func updateTaskStatus() {
	tasksExecutions := dao.GetExecutionsToWatch()
	for _, tasksExecution := range tasksExecutions {
		status, err := k8s_client.GetJobStatus(tasksExecution.ID)
		if err != nil {
			log.Fatal("Error while gettign job status ", err)
		}
		if status != tasksExecution.Status {
			log.Println("Updating task status ", tasksExecution.ID)
			dao.UpdateExecutionStatus(tasksExecution.ID, status)
		}
	}
}
