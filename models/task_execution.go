package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TaskExecution struct {
	gorm.Model
	ID                      uuid.UUID      `json:"id" gorm:"type:uuid` //;default:uuid_generate_v4()"
	TaskDefinitionId        uuid.UUID      `json:"task_definition_id"`
	Status                  Status         `json:"status"`
	Image                   string         `json:"image" binding:"required"`
	Name                    string         `json:"name"`
	TtlSecondsAfterFinished int            `json:"ttl_seconds_after_finished"`
	NameSpace               string         `json:"namespace"`
	Args                    string         `json:"args"`
	Cmd                     pq.StringArray `json:"cmd" gorm:"type:text[]"`
	Metadata                datatypes.JSON `json:"metadata"`
	ExecutionParameters     datatypes.JSON `json:"execution_parameters"`
}
