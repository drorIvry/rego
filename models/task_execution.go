package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TaskExecution struct {
	gorm.Model
	Status                  Status
	TaskDefinitionId        uint           `json:"task_definition_id"`
	Image                   string         `json:"image" binding:"required"`
	Name                    string         `json:"name"`
	TtlSecondsAfterFinished int            `json:"ttl_seconds_after_finished"`
	NameSpace               string         `json:"nameSpace"`
	ExecutionInterval       int            `json:"execution_interval"`
	NextExecutionTime       time.Time      `json:"next_execution_time"`
	Enabled                 bool           `json:"enabled"`
	Deleted                 bool           `json:"deleted"`
	Args                    string         `json:"args"`
	Cmd                     string         `json:"cmd"`
	Metadata                datatypes.JSON `json:"metadata"`
	ExecutionParameters     datatypes.JSON `json:"execution_parameters"`
}
