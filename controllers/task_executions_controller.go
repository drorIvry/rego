package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"

	dao "github.com/drorivry/matter/dao"
	k8s_client "github.com/drorivry/matter/k8s"
	"github.com/drorivry/matter/models"
	"github.com/gin-gonic/gin"
)

// AbortTaskExecution             godoc
// @Summary      				  Abort a running task and kill the pod
// @Description                   Kill a running k8s job and update its task execution
// @Tags                          execution
// @Produce                       json
// @Param                         executionId  path string  true  "The task execution id"
// @Success                       200
// @Router                        /execution/{executionId}/abort [get]
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

	dao.UpdateExecutionStatus(executionId, models.ABORTED)

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "aborted",
		},
	)
}
