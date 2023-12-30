package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExecutionStatusHistory struct {
	gorm.Model
	ExecutionID uuid.UUID `json:"execution_id,omitempty" gorm:"type:uuid`
}
