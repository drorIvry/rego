package dao

import (
	"log"
	"time"

	"github.com/drorivry/matter/config"
	"github.com/drorivry/matter/initializers"
	"github.com/drorivry/matter/models"
	"github.com/google/uuid"
)

func GetTasksToTimeout() []models.TaskDefinition {
	var tasks []models.TaskDefinition
	timeoutTime := time.Now().Add(time.Duration(-config.TASK_TIMEOUT) * time.Second)
	initializers.GetTaskExecutionsTable().Where("status < 400 AND created_at < ?", timeoutTime).Find(&tasks)
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

func UpdateExecutionAborted(executionId uuid.UUID, status models.Status) {
	initializers.GetTaskExecutionsTable().Where(
		"id = ?",
		executionId,
	).Updates(
		models.TaskExecution{
			Status: status,
		},
	)
}
