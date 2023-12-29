package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TaskExecution struct {
	gorm.Model
	ID               uuid.UUID      `json:"id,omitempty" gorm:"type:uuid` //;default:uuid_generate_v4()"
	CreationTime     time.Time      `json:"created_at"`
	LastModified     time.Time      `json:"last_modified"`
	TaskDefinitionId uuid.UUID      `json:"task_definition_id,omitempty"`
	StatusCode       Status         `json:"status_code,omitempty"`
	TaskStatus       string         `json:"status,omitempty"`
	Image            string         `json:"image,omitempty" binding:"required"`
	Name             string         `json:"name,omitempty"`
	Namespace        string         `json:"namespace,omitempty"`
	Cmd              pq.StringArray `json:"cmd,omitempty" gorm:"type:text[]"`
	Metadata         datatypes.JSON `json:"metadata,omitempty"`
}
