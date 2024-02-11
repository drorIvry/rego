package poller

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/drorivry/rego/dao"
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
	tasks := dao.InternalGetPendingTasks()
	for _, task := range tasks {
		log.Info().Str("defintion_id", task.ID.String()).Msg("deploying task")
		DeployJob(&task)
	}
}

func DeployJob(task *models.TaskDefinition) {
	taskEx := models.CreateExecutionFromDefinition(task)
	jobName := k8s_client.BuildJobName(&taskEx)

	taskEx.StatusCode = models.JOB_DEPLOYED

	dao.InsertTaskExecution(&taskEx)

	task.ExecutionsCounter++

	// TODO: Move to dao
	if task.ExecutionInterval > 0 {
		task.NextExecutionTime = time.Now().Add(time.Duration(task.ExecutionInterval) * time.Second)
	} else {
		task.Enabled = false
	}
	dao.UpdateDefinition(task)
	dao.UpdateExecutionStatus(taskEx.ID, models.JOB_DEPLOYED, task.OrganizationId)
	k8s_client.LaunchK8sJob(&jobName, &taskEx)
}

func timeoutTasks() {
	taskExecutions := dao.GetTasksToTimeout()

	for _, taskExecution := range taskExecutions {
		log.Info().Str(
			"execution_id",
			taskExecution.ID.String(),
		).Msg("timing out task")

		err := k8s_client.AbortTask(taskExecution.ID)
		if err != nil {
			log.Error().Err(err).Msg("Could not abort task")
			continue
		}
		dao.UpdateExecutionStatus(taskExecution.ID, models.TIMEOUT, taskExecution.OrganizationId)
	}
}

func updateTaskStatus() {
	taskExecutions := dao.GetExecutionsToWatch()
	for _, taskExecution := range taskExecutions {
		status, err := k8s_client.GetJobStatus(taskExecution.ID)
		if err != nil {
			log.Error().Err(
				err,
			).Str(
				"execution_id",
				taskExecution.ID.String(),
			).Msg(
				"Error while getting job status",
			)
		}
		if status != taskExecution.StatusCode {
			log.Info().Str(
				"execution_id",
				taskExecution.ID.String(),
			).Str(
				"status",
				models.NumericStatusToStringStatus(status),
			).Msg("Updating task status")
			dao.UpdateExecutionStatus(taskExecution.ID, status, taskExecution.OrganizationId)
			dao.UpdateDefinitionStatus(taskExecution.TaskDefinitionId, status, taskExecution.OrganizationId)
		}
	}
}
