package models

import "gorm.io/gorm"

type TaskDefinition struct {
	gorm.Model
	image string
	ttlSecondsAfterFinished int
}
