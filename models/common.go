package models

import (
	"time"

	"github.com/google/uuid"
)

type Status int

const (
	READY        Status = 100
	JOB_DEPLOYED Status = 200
	RUNNING      Status = 300
	TIMEOUT      Status = 400
	PROC_ERROR   Status = 410
	APP_ERROR    Status = 420
	ABORTED      Status = 430
	SUCCESS      Status = 500
)

func CreateExecutionFromDefinition(taskdef TaskDefinition) TaskExecution {
	taskEx := TaskExecution{
		ID:                      uuid.New(),
		Status:                  READY,
		TaskDefinitionId:        taskdef.ID,
		Image:                   taskdef.Image,
		Name:                    taskdef.Name,
		NameSpace:               taskdef.NameSpace,
		TtlSecondsAfterFinished: taskdef.TtlSecondsAfterFinished,
		Args:                    taskdef.Args,
		Cmd:                     taskdef.Cmd,
		Metadata:                taskdef.Metadata,
		ExecutionParameters:     taskdef.ExecutionParameters,
	}

	taskEx.CreatedAt = time.Now()

	return taskEx
}
