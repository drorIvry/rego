package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const TASK_EXECUTIONS_TABLE_NAME string = "rego_task_executions"

type TaskExecution struct {
	gorm.Model
	ID               uuid.UUID      `json:"id,omitempty"` //;default:uuid_generate_v4()"
	OrganizationId   string         `json:"organization_id,omitempty"`
	TaskDefinitionId uuid.UUID      `json:"task_definition_id,omitempty"`
	StatusCode       Status         `json:"status_code,omitempty"`
	TaskStatus       string         `json:"status,omitempty"`
	Image            string         `json:"image,omitempty" binding:"required"`
	Name             string         `json:"name,omitempty"`
	Namespace        string         `json:"namespace,omitempty"`
	Cmd              StringArray    `json:"cmd"`
	Metadata         datatypes.JSON `json:"metadata,omitempty"`
}

func (TaskExecution) TableName() string {
	return TASK_EXECUTIONS_TABLE_NAME
}
