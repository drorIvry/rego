package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TaskDefinition struct {
	gorm.Model
	Image                   string         `json:"image" binding:"required"`
	Name                    string         `json:"name"`
	NameSpace               string         `json:"nameSpace"`
	TtlSecondsAfterFinished int            `json:"ttl_seconds_after_finished"`
	ExecutionInterval       int            `json:"execution_interval"`
	ExecutionsCounter       int            `json:"execution_counter"`
	NextExecutionTime       time.Time      `json:"next_execution_time"`
	Enabled                 bool           `json:"enabled"`
	Deleted                 bool           `json:"deleted"`
	Args                    string         `json:"args"`
	Cmd                     string         `json:"cmd"`
	Metadata                datatypes.JSON `json:"metadata"`
	ExecutionParameters     datatypes.JSON `json:"execution_parameters"`
}
