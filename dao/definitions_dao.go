package dao

import (
	"log"
	"time"

	"github.com/drorivry/rego/initializers"
	"github.com/drorivry/rego/models"
	"github.com/google/uuid"
)

func GetPendingTasks() []models.TaskDefinition {
	var tasks []models.TaskDefinition
	initializers.GetTaskDefinitionsTable().Where(
		"enabled = ?",
		true,
	).Where(
		"next_execution_time < ?",
		time.Now(),
	).Find(&tasks)
	return tasks
}

func initDefaultFields(taskDef *models.TaskDefinition) {
	taskDef.CreationTime = time.Now()
	taskDef.LastModified = time.Now()
	taskDef.Deleted = false
	taskDef.Enabled = true
	taskDef.ExecutionsCounter = 0

	if taskDef.Namespace == "" {
		taskDef.Namespace = "default"
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

func GetTaskDefinitionById(definitionId uuid.UUID) (models.TaskDefinition, error) {
	var task_def models.TaskDefinition
	result := initializers.GetTaskDefinitionsTable().Where(
		"id = ?",
		definitionId,
	).First(
		&task_def,
	)

	return task_def, result.Error
}

func UpdateDefinition(taskDefinition models.TaskDefinition) {
	taskDefinition.LastModified = time.Now()
	initializers.GetTaskDefinitionsTable().Where(
		"id = ?",
		taskDefinition.ID,
	).Updates(
		taskDefinition,
	)
}

func DeleteTaskDefinitionById(definitionId uuid.UUID) {
	initializers.GetTaskDefinitionsTable().Where(
		"id = ?",
		definitionId,
	).Updates(
		models.TaskDefinition{
			Deleted:      true,
			Enabled:      false,
			LastModified: time.Now(),
		},
	).Delete(
		"id = ?",
		definitionId,
	)
}
