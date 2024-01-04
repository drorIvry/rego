package dao

import (
	"github.com/drorivry/rego/initializers"
	"github.com/drorivry/rego/models"
	"github.com/google/uuid"
)

func GetExecutionStatusHistory(executionId uuid.UUID) []models.ExecutionStatusHistory {
	var statusHistory []models.ExecutionStatusHistory

	initializers.GetExecutionsStatusHistoryTable().Where(
		"execution_id = ?",
		executionId.String(),
	).Order(
		"created_at DESC",
	).Find(
		&statusHistory,
	)
	return statusHistory
}

func InsertExecutionStatusUpdate(executionId uuid.UUID, status models.Status) {
	model := models.ExecutionStatusHistory{
		ExecutionID: executionId,
		Status:      status,
	}
	initializers.GetExecutionsStatusHistoryTable().Create(&model)
}