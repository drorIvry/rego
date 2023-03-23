package models

type Status int

const (
	READY        Status = 100
	JOB_DEPLOYED Status = 200
	RUNNING      Status = 300
	TIMEOUT      Status = 400
	PROC_ERROR   Status = 410
	APP_ERROR    Status = 420
	SUCCESS      Status = 500
)

func CreateExecutionFromDefinition(taskdef TaskDefinition) TaskExecution {
	return TaskExecution{
		Status:                  READY,
		TaskDefinitionId:        taskdef.ID,
		Image:                   taskdef.Image,
		Name:                    taskdef.Name,
		NameSpace:               taskdef.NameSpace,
		TtlSecondsAfterFinished: taskdef.TtlSecondsAfterFinished,
		ExecutionInterval:       taskdef.ExecutionInterval,
		NextExecutionTime:       taskdef.NextExecutionTime,
		Enabled:                 taskdef.Enabled,
		Deleted:                 taskdef.Deleted,
		Args:                    taskdef.Args,
		Cmd:                     taskdef.Cmd,
		Metadata:                taskdef.Metadata,
		ExecutionParameters:     taskdef.ExecutionParameters,
	}
}
