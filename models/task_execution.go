package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TaskExecution struct {
	gorm.Model
	ID                      uuid.UUID      `json:"id,omitempty" gorm:"type:uuid` //;default:uuid_generate_v4()"
	TaskDefinitionId        uuid.UUID      `json:"task_definition_id,omitempty"`
	StatusCode              Status         `json:"status_code,omitempty"`
	TaskStatus              string         `json:"status,omitempty"`
	Image                   string         `json:"image,omitempty" binding:"required"`
	Name                    string         `json:"name,omitempty"`
	TtlSecondsAfterFinished int            `json:"ttl_seconds_after_finished,omitempty"`
	Namespace               string         `json:"namespace,omitempty"`
	Args                    string         `json:"args,omitempty"`
	Cmd                     pq.StringArray `json:"cmd,omitempty" gorm:"type:text[]"`
	Metadata                datatypes.JSON `json:"metadata,omitempty"`
	ExecutionParameters     datatypes.JSON `json:"execution_parameters,omitempty"`
}
