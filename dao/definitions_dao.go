package dao

import (
	"time"

	"github.com/rs/zerolog/log"

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
		log.Error().Err(result.Error).Msg("Error saving to database")
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
	initializers.GetTaskDefinitionsTable().Where(
		"id = ?",
		taskDefinition.ID,
	).Where(
		"deleted = ?", // Can't update deleted definitions
		false,
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
			Deleted: true,
			Enabled: false,
		},
	).Delete(
		"id = ?",
		definitionId,
	)
}
