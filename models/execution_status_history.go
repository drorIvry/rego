package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const EXECUTION_STATUS_HISTORY_TABLE_NAME string = "rego_execution_status_history"

type ExecutionStatusHistory struct {
	gorm.Model
	ExecutionID    uuid.UUID `json:"execution_id,omitempty"`
	OrganizationId string    `json:"organization_id,omitempty"`
	TaskStatus     string    `json:"status,omitempty"`
}

func (ExecutionStatusHistory) TableName() string {
	return EXECUTION_STATUS_HISTORY_TABLE_NAME
}
