package poller

import (
	"log"
	"time"

	"github.com/drorivry/rego/dao"
	"github.com/drorivry/rego/initializers"
	"github.com/drorivry/rego/models"

	k8s_client "github.com/drorivry/rego/k8s"
)

type Poller struct {
	interval int
	quit     chan struct{}
}

func NewPoller(interval int) *Poller {
	p := Poller{
		interval: interval,
		quit:     make(chan struct{}),
	}

	return &p
}

func (p *Poller) Run() {
	ticker := time.NewTicker(time.Duration(p.interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			go deployReadyTasks()
			go timeoutTasks()
			go updateTaskStatus()
		case <-p.quit:
			ticker.Stop()
			return
		}
	}
}

func (p *Poller) Shutdown() {
	close(p.quit)
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

	taskEx.StatusCode = models.JOB_DEPLOYED
	taskEx.TaskStatus = models.NumericStatusToStringStatus(models.JOB_DEPLOYED)

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
			log.Println("Error while getting job status ", err)
		}
		if status != tasksExecution.StatusCode {
			log.Println("Updating task status ", tasksExecution.ID)
			dao.UpdateExecutionStatus(tasksExecution.ID, status)
		}
	}
}
