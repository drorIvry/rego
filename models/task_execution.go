package models

type TaskExecution struct {
	TaskDefinition
	executionParameter map[string]string
}
