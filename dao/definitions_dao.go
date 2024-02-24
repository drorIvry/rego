package dao

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/drorivry/rego/initializers"
	"github.com/drorivry/rego/models"
	"github.com/google/uuid"
)

func GetPendingTasks(OrganizationId string) []models.TaskDefinition {
	var tasks []models.TaskDefinition
	initializers.GetTaskDefinitionsTable().Where(
		"enabled = ?",
		true,
	).Where(
		"organization_id = ?",
		OrganizationId,
	).Where(
		"deleted = ?",
		false,
	).Where(
		"next_execution_time < ?",
		time.Now(),
	).Find(&tasks)
	return tasks
}

func InternalGetPendingTasks() []models.TaskDefinition {
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

func GetAllTaskDefinitions(OrganizationId string, offset int, limit int) []models.TaskDefinition {
	var tasks []models.TaskDefinition
	initializers.GetTaskDefinitionsTable().Where(
		"organization_id = ?",
		OrganizationId,
	).Order(
		"created_at desc",
	).Offset(
		offset,
	).Limit(
		limit,
	).Find(
		&tasks,
	)
	return tasks
}

func GetTaskDefinitionById(definitionId uuid.UUID, OrganizationId string) (models.TaskDefinition, error) {
	var task_def models.TaskDefinition
	result := initializers.GetTaskDefinitionsTable().Where(
		"id = ?",
		definitionId,
	).Where(
		"organization_id = ?",
		OrganizationId,
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
		"organization_id = ?",
		taskDefinition.OrganizationId,
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

func DeleteTaskDefinitionById(definitionId uuid.UUID, OrganizationId string) {
	result := initializers.GetTaskDefinitionsTable().Select(
		"*", // Selecting to update all even 0 value columns
	).Where(
		"id = ?",
		definitionId,
	).Where(
		"organization_id = ?",
		OrganizationId,
	).Delete(
		&models.TaskDefinition{},
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

func UpdateDefinitionStatus(definitionId uuid.UUID, status models.Status, OrganizationId string) {
	result := initializers.GetTaskDefinitionsTable().Where(
		"id = ?",
		definitionId,
	).Where(
		"organization_id = ?",
		OrganizationId,
	).Updates(
		models.TaskDefinition{
			LatestStatus: models.NumericStatusToStringStatus(status),
		},
	)

	if result.Error != nil {
		log.Error().Err(
			result.Error,
		).Str(
			"definition_id",
			definitionId.String(),
		).Msg(
			"Couldn't update definition status",
		)
		return
	}
}
