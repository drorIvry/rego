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
