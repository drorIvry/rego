package dao

import (
	"github.com/drorivry/rego/initializers"
	"github.com/drorivry/rego/models"
	"github.com/google/uuid"
)

func GetExecutionStatusHistory(executionId uuid.UUID, OrganizationId string ) []models.ExecutionStatusHistory {
	var statusHistory []models.ExecutionStatusHistory

	initializers.GetExecutionsStatusHistoryTable().Where(
		"organization_id = ?",
		OrganizationId,
	).Where(
		"execution_id = ?",
		executionId.String(),
	).Order(
		"created_at DESC",
	).Find(
		&statusHistory,
	)
	return statusHistory
}

func InsertExecutionStatusUpdate(executionId uuid.UUID, status models.Status, OrganizationId string) {
	model := models.ExecutionStatusHistory{
		ExecutionID: executionId,
		TaskStatus:  models.NumericStatusToStringStatus(status),
		OrganizationId: OrganizationId,
	}
	initializers.GetExecutionsStatusHistoryTable().Create(&model)
}
