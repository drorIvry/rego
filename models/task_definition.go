package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const TASK_DEFINITIONS_TABLE_NAME string = "rego_task_definitions"

type TaskDefinition struct {
	gorm.Model
	ID                uuid.UUID      `json:"id"` //;default:uuid_generate_v4()"
	OrganizationId    string         `json:"organization_id,omitempty"`
	Image             string         `json:"image" binding:"required"`
	Name              string         `json:"name"`
	Namespace         string         `json:"namespace"`
	ExecutionInterval int            `json:"execution_interval"`
	ExecutionsCounter int            `json:"execution_counter"`
	NextExecutionTime time.Time      `json:"next_execution_time"`
	Enabled           bool           `json:"enabled"`
	Deleted           bool           `json:"deleted"`
	Cmd               StringArray    `json:"cmd"`
	Metadata          datatypes.JSON `json:"metadata"`
}

func (TaskDefinition) TableName() string {
	return TASK_DEFINITIONS_TABLE_NAME
}
