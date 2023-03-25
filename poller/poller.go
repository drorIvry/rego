package poller

import (
	"log"
	"time"

	"github.com/drorivry/matter/dao"
	"github.com/drorivry/matter/initializers"
	"github.com/drorivry/matter/models"
	"k8s.io/client-go/kubernetes"

	k8s_client "github.com/drorivry/matter/k8s"
)

func Run(interval int, clientset *kubernetes.Clientset) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			go deployReadyTasks(clientset)
			go timeoutTasks()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func FormatInt(u uint) {
	panic("unimplemented")
}

func deployReadyTasks(clientset *kubernetes.Clientset) {
	tasks := dao.GetPendingTasks()

	for _, task := range tasks {
		log.Println("deploying task ", task.ID)
		taskEx := models.CreateExecutionFromDefinition(task)
		jobName := k8s_client.BuildJobName(taskEx)
		k8s_client.LaunchK8sJob(clientset, &jobName, &taskEx)

		taskEx.Status = models.JOB_DEPLOYED

		dao.InsertTaskExecution(taskEx)

		task.ExecutionsCounter++

		if task.ExecutionInterval > 0 {
			task.NextExecutionTime = time.Now().Add(time.Duration(task.ExecutionInterval) * time.Second)
		} else {
			task.Enabled = false
		}

		initializers.DefinitionsTable.Save(&task)
		initializers.ExecutionsTable.Save(&taskEx)
	}
}

func timeoutTasks() {
	//log.Println("killing timed out tasks")
}
