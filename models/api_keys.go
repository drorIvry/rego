package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const API_KEYS_TABLE_NAME string = "api_keys"

type ApiKeys struct {
	gorm.Model
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid`
	ApiKey string `json:"api_key,omitempty"`
	ApiKeyHint string `json:"api_key_hint,omitempty"`
	OrganizationId string `json:"organization_id,omitempty"`
	TaskStatus  string    `json:"status,omitempty"`
}

func (ApiKeys) TableName() string {
	return API_KEYS_TABLE_NAME
}
