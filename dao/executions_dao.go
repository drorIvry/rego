package dao

import (
	"log"
	"time"

	"github.com/drorivry/matter/config"
	"github.com/drorivry/matter/initializers"
	"github.com/drorivry/matter/models"
	"github.com/google/uuid"
)

func GetTasksToTimeout() []models.TaskExecution {
	var tasks []models.TaskExecution
	timeoutTime := time.Now().Add(time.Duration(-config.TASK_TIMEOUT) * time.Second)
	initializers.GetTaskExecutionsTable().Where("status_code < 400 AND created_at < ?", timeoutTime).Find(&tasks)
	return tasks
}

func CreateTaskExecution(taskDef models.TaskExecution) error {
	result := initializers.DB.Create(&taskDef)
	if result.Error != nil {
		log.Panic("Error saving to database")
		return result.Error
	}
	return nil
}

func InsertTaskExecution(taskEx models.TaskExecution) error {
	result := initializers.GetTaskExecutionsTable().Create(&taskEx)
	if result.Error != nil {
		log.Panic("Error saving to database", result.Error)
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
	initializers.GetTaskExecutionsTable().Where("status_code < ?", models.TIMEOUT).Find(&tasks)
	return tasks
}

func GetLatestExecutionByDefinitionId(definitionId uuid.UUID) models.TaskExecution {
	var task models.TaskExecution
	initializers.GetTaskExecutionsTable().Where("task_definition_id = ?", definitionId).Order("created_at desc").First(&task)
	return task
}
