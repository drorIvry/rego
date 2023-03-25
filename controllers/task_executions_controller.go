package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"

	dao "github.com/drorivry/matter/dao"
	k8s_client "github.com/drorivry/matter/k8s"
	"github.com/gin-gonic/gin"
)

func AbortTaskExecution(c *gin.Context) {
	var executionIdUnparsed = strings.TrimSpace(c.Param("executionId"))

	log.Println("Aborting task", executionIdUnparsed)

	var executionId, err = uuid.Parse(executionIdUnparsed)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = k8s_client.AbortTask(executionId)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dao.UpdateExecutionAborted(executionId)

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "aborted",
		},
	)
}
