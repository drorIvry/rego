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
		"deleted = ?",
		false,
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
	taskDef.NextExecutionTime = time.Now()

	if taskDef.Namespace == "" {
		taskDef.Namespace = "default"
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

	if result.Error != nil {
		log.Error().Err(
			result.Error,
		).Str(
			"definition_id",
			definitionId.String(),
		).Msg(
			"Couldn't get task definition",
		)
	}

	return task_def, result.Error
}

func UpdateDefinition(taskDefinition *models.TaskDefinition) {
	result := initializers.GetTaskDefinitionsTable().Where(
		"id = ?",
		taskDefinition.ID,
	).Where(
		"deleted = ?", // Can't update deleted definitions
		false,
	).Save(
		taskDefinition,
	)

	if result.Error != nil {
		log.Error().Err(
			result.Error,
		).Str(
			"definition_id",
			taskDefinition.ID.String(),
		).Msg(
			"Couldn't update task definition",
		)
	}
}

func DeleteTaskDefinitionById(definitionId uuid.UUID) {
	result := initializers.GetTaskDefinitionsTable().Select(
		"*", // Selecting to update all even 0 value columns
	).Where(
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

	if result.Error != nil {
		log.Error().Err(
			result.Error,
		).Str(
			"definition_id",
			definitionId.String(),
		).Msg(
			"Couldn't update task definition",
		)
	}
}
