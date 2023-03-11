package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TaskDefinition struct {
	gorm.Model
	Image                   string
	TtlSecondsAfterFinished int
	Status                  Status `gorm:"type:text"`
	ExecutionInterval       int
	ExecutionsCounter       int
	NextExecutionTime       time.Time
	Enabled                 bool
	Deleted                 bool
	Args                    []string `gorm:"type:text"`
	Cmd                     string
	Metadata                datatypes.JSON
}
