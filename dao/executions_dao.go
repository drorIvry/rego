package dao

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/drorivry/rego/config"
	"github.com/drorivry/rego/initializers"
	"github.com/drorivry/rego/models"
	"github.com/google/uuid"
)

func GetTasksToTimeout() []models.TaskExecution {
	var executions []models.TaskExecution
	timeoutTime := time.Now().Add(time.Duration(-config.TASK_TIMEOUT) * time.Second)
	initializers.GetTaskExecutionsTable().Where(
		"status_code < ?",
		400,
	).Where(
		"created_at < ?",
		timeoutTime,
	).Find(&executions)
	return executions
}

func InsertTaskExecution(taskEx models.TaskExecution) error {
	result := initializers.GetTaskExecutionsTable().Create(&taskEx)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Error saving to database")
		return result.Error
	}
	return nil
}

func GetExecutionById(executionId uuid.UUID) *models.TaskExecution {
	var execution models.TaskExecution
	initializers.GetTaskExecutionsTable().Where(
		"id = ?",
		executionId,
	).First(
		&execution,
	)

	return &execution
}

func UpdateExecutionStatus(executionId uuid.UUID, status models.Status) {
	initializers.GetTaskExecutionsTable().Where(
		"id = ?",
		executionId,
	).Updates(
		models.TaskExecution{
			StatusCode: status,
			TaskStatus: models.NumericStatusToStringStatus(status),
		},
	)
}

func GetExecutionsToWatch() []models.TaskExecution {
	var tasks []models.TaskExecution
	initializers.GetTaskExecutionsTable().Where(
		"status_code < ?",
		models.TIMEOUT,
	).Find(&tasks)
	return tasks
}

func GetLatestExecutionByDefinitionId(definitionId uuid.UUID) (models.TaskExecution, error) {
	var task models.TaskExecution
	result := initializers.GetTaskExecutionsTable().Where(
		"task_definition_id = ?",
		definitionId,
	).Order(
		"created_at desc",
	).First(&task)

	return task, result.Error
}
