package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const EXECUTION_STATUS_HISTORY_TABLE_NAME string = "execution_status_history"

type ExecutionStatusHistory struct {
	gorm.Model
	ExecutionID uuid.UUID `json:"execution_id,omitempty" gorm:"type:uuid`
	Status      Status    `json:"status" gorm:"type:text`
}

func (ExecutionStatusHistory) TableName() string {
	return EXECUTION_STATUS_HISTORY_TABLE_NAME
}
