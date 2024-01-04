package models

import (
	"time"

	"github.com/google/uuid"
)

type Status int

const (
	READY        Status = 100
	JOB_DEPLOYED Status = 200
	PENDING      Status = 250
	RUNNING      Status = 300
	TIMEOUT      Status = 400
	PROC_ERROR   Status = 410
	APP_ERROR    Status = 420
	ABORTED      Status = 430
	SUCCESS      Status = 500
)

func (s *Status) String() string {
	if *s == READY {
		return "READY"
	} else if *s == JOB_DEPLOYED {
		return "JOB_DEPLOYED"
	} else if *s == PENDING {
		return "PENDING"
	} else if *s == RUNNING {
		return "RUNNING"
	} else if *s == TIMEOUT {
		return "TIMEOUT"
	} else if *s == PROC_ERROR {
		return "PROC_ERROR"
	} else if *s == APP_ERROR {
		return "APP_ERROR"
	} else if *s == ABORTED {
		return "ABORTED"
	} else if *s == SUCCESS {
		return "SUCCESS"
	} else {
		return "UNKNOWN"
	}
}

func NumericStatusToStringStatus(status Status) string {
	switch status {
	case READY:
		return "READY"
	case JOB_DEPLOYED:
		return "JOB_DEPLOYED"
	case PENDING:
		return "PENDING"
	case RUNNING:
		return "RUNNING"
	case TIMEOUT:
		return "TIMEOUT"
	case PROC_ERROR:
		return "PROC_ERROR"
	case APP_ERROR:
		return "APP_ERROR"
	case ABORTED:
		return "ABORTED"
	case SUCCESS:
		return "SUCCESS"
	default:
		return "UNKNOWN"
	}
}

func CreateExecutionFromDefinition(taskDef TaskDefinition) TaskExecution {
	taskEx := TaskExecution{
		ID:               uuid.New(),
		StatusCode:       READY,
		TaskStatus:       NumericStatusToStringStatus(READY),
		TaskDefinitionId: taskDef.ID,
		Image:            taskDef.Image,
		Name:             taskDef.Name,
		Namespace:        taskDef.Namespace,
		Cmd:              taskDef.Cmd,
		Metadata:         taskDef.Metadata,
	}

	taskEx.CreatedAt = time.Now()

	return taskEx
}
