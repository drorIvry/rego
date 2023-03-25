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
	ID                      uuid.UUID      `json:"id" gorm:"type:uuid` //;default:uuid_generate_v4()"
	Image                   string         `json:"image" binding:"required"`
	Name                    string         `json:"name"`
	NameSpace               string         `json:"namespace"`
	TtlSecondsAfterFinished int            `json:"ttl_seconds_after_finished"`
	ExecutionInterval       int            `json:"execution_interval"`
	ExecutionsCounter       int            `json:"execution_counter"`
	NextExecutionTime       time.Time      `json:"next_execution_time"`
	Enabled                 bool           `json:"enabled"`
	Deleted                 bool           `json:"deleted"`
	Args                    string         `json:"args"`
	Cmd                     pq.StringArray `json:"cmd" gorm:"type:text[]"`
	Metadata                datatypes.JSON `json:"metadata"`
	ExecutionParameters     datatypes.JSON `json:"execution_parameters"`
}
