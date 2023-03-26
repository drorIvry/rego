package dao

import (
	"log"
	"time"

	"github.com/drorivry/matter/initializers"
	"github.com/drorivry/matter/models"
	"github.com/google/uuid"
)

func GetPendingTasks() []models.TaskDefinition {
	var tasks []models.TaskDefinition
	initializers.GetTaskDefinitionsTable().Where("enabled = true AND next_execution_time < ?", time.Now()).Find(&tasks)
	return tasks
}

func initDefaultFields(taskDef *models.TaskDefinition) {
	taskDef.CreatedAt = time.Now()
	taskDef.Deleted = false
	taskDef.Enabled = true
	taskDef.ExecutionsCounter = 0

	if taskDef.NameSpace == "" {
		taskDef.NameSpace = "default"
	}

	if taskDef.ExecutionInterval > 0 {
		taskDef.NextExecutionTime = time.Now().Add(time.Duration(taskDef.ExecutionInterval) * time.Second)
	}
}

func CreateTaskDefinition(taskDef *models.TaskDefinition) error {
	initDefaultFields(taskDef)

	result := initializers.GetTaskDefinitionsTable().Create(taskDef)
	if result.Error != nil {
		log.Panic("Error saving to database", result.Error)
		return result.Error
	}
	return nil
}


func GetAllTaskDefinitions() []models.TaskDefinition {
	var tasks []models.TaskDefinition
	initializers.GetTaskDefinitionsTable().Find(&tasks)
	return tasks
}


func GetTaskDefinitionById(definitionId uuid.UUID) models.TaskDefinition {
	var task_def models.TaskDefinition
	initializers.GetTaskDefinitionsTable().Where(
		"id = ?",
		definitionId,
	).First(
		&task_def,
	)

	return task_def
}

