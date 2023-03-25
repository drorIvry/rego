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
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func deployReadyTasks() {
	tasks := dao.GetPendingTasks()
	log.Println("polling")

	for _, task := range tasks {
		log.Println("deploying task ", task.ID)
		taskEx := models.CreateExecutionFromDefinition(task)
		jobName := k8s_client.BuildJobName(taskEx)
		k8s_client.LaunchK8sJob(clientset, &jobName, &taskEx)
		DeployJob(task)
	}
}

func DeployJob(task models.TaskDefinition) {
	taskEx := models.CreateExecutionFromDefinition(task)
	jobName := k8s_client.BuildJobName(taskEx)

	taskEx.Status = models.JOB_DEPLOYED

	dao.InsertTaskExecution(taskEx)

	task.ExecutionsCounter++

	initializers.DefinitionsTable.Save(&task)
	initializers.ExecutionsTable.Save(&taskEx)
	if task.ExecutionInterval > 0 {
		task.NextExecutionTime = time.Now().Add(time.Duration(task.ExecutionInterval) * time.Second)
	} else {
		task.Enabled = false
	}

	initializers.DB.Table("task_definitions").Save(&task)
	initializers.DB.Table("task_executions").Save(&taskEx)
	k8s_client.LaunchK8sJob(&jobName, &taskEx)
}

func timeoutTasks() {
	//log.Println("killing timed out tasks")
}
