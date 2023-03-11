package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TaskExecution struct {
	gorm.Model
	executionParameter      map[string]string
	Status                  Status
	TaskDefinitionId        uint
	Image                   string
	TtlSecondsAfterFinished int
	NextExecutionTime       time.Time
	Enabled                 bool
	Deleted                 bool
	Args                    []string `gorm:"type:text"`
	Cmd                     string
	Metadata                datatypes.JSON
}
