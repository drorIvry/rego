package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TaskDefinition struct {
	gorm.Model
	ID                      uuid.UUID      `json:"id,omitempty" gorm:"type:uuid` //;default:uuid_generate_v4()"
	Image                   string         `json:"image,omitempty" binding:"required"`
	Name                    string         `json:"name,omitempty"`
	Namespace               string         `json:"namespace,omitempty"`
	TtlSecondsAfterFinished int            `json:"ttl_seconds_after_finished,omitempty"`
	ExecutionInterval       int            `json:"execution_interval,omitempty"`
	ExecutionsCounter       int            `json:"execution_counter,omitempty"`
	NextExecutionTime       time.Time      `json:"next_execution_time,omitempty"`
	Enabled                 bool           `json:"enabled,omitempty"`
	Deleted                 bool           `json:"deleted,omitempty"`
	Args                    string         `json:"args,omitempty"`
	Cmd                     pq.StringArray `json:"cmd,omitempty" gorm:"type:text[]"`
	Metadata                datatypes.JSON `json:"metadata,omitempty"`
	ExecutionParameters     datatypes.JSON `json:"execution_parameters,omitempty"`
}
