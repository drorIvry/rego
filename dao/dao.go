package dao

import (
	"log"
	"time"

	"github.com/drorivry/matter/initializers"
	"github.com/drorivry/matter/models"
)

func GetPendingTasks() []models.TaskDefinition {
	var tasks []models.TaskDefinition
	initializers.DB.Table("task_definitions").Where("enabled = true AND next_execution_time < ?", time.Now()).Find(&tasks)
	return tasks
}

func initDefaultFields(taskDef *models.TaskDefinition) {
	taskDef.CreatedAt = time.Now()
	taskDef.Deleted = false
	taskDef.Enabled = true
	taskDef.ExecutionsCounter = 0

	if taskDef.ExecutionInterval > 0 {
		taskDef.NextExecutionTime = time.Now().Add(time.Duration(taskDef.ExecutionInterval) * time.Second)
	}
}

func CreateTaskDefinition(taskDef *models.TaskDefinition) error {
	initDefaultFields(taskDef)

	result := initializers.DB.Table("task_definitions").Create(taskDef)
	if result.Error != nil {
		log.Panic("Error saving to database", result.Error)
		return result.Error
	}
	return nil
}

func CreateTaskExecution(taskDef models.TaskExecution) error {
	result := initializers.DB.Create(&taskDef)
	if result.Error != nil {
		log.Panic("Error saving to database")
		return result.Error
	}
	return nil
}

func GetAllTaskDefinitions() []models.TaskDefinition {
	var tasks []models.TaskDefinition
	initializers.DB.Table("task_definitions").Find(&tasks)
	return tasks
}

func InsertTaskExecution(taskEx models.TaskExecution) error {
	result := initializers.DB.Table("task_executions").Create(&taskEx)
	if result.Error != nil {
		log.Panic("Error saving to database", result.Error)
		return result.Error
	}
	return nil
}
