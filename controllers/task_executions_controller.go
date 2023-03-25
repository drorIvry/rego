package controllers

import (
	"strconv"

	k8s_client "github.com/drorivry/matter/k8s"
	"github.com/gin-gonic/gin"
)

func AbortTaskExecution(c *gin.Context) {
	var execIdString = c.Param("execution_id")
	executionId, err := strconv.Atoi(execIdString)

	k8s_client.AbortTask(executionId)

}
