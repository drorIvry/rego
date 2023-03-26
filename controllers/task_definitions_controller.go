package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/drorivry/matter/dao"
	"github.com/drorivry/matter/models"
	"github.com/drorivry/matter/poller"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func CreateTaskDefinition(c *gin.Context) {
	var newTaskDef models.TaskDefinition

	if err := c.BindJSON(&newTaskDef); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newTaskDef.ID = uuid.New()

	err := dao.CreateTaskDefinition(&newTaskDef)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"message":       "created",
			"definition_id": newTaskDef.ID.String(),
		},
	)
}

func GetAllTaskDefinitions(c *gin.Context) {
	tasks := dao.GetAllTaskDefinitions()
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetAllPendingTaskDefinitions(c *gin.Context) {
	tasks := dao.GetPendingTasks()
	c.IndentedJSON(http.StatusOK, tasks)
}

func RerunTask(c *gin.Context) {
	uuidParam := c.Param("definitionId")
	var definitionId, err = uuid.Parse(uuidParam)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	task := dao.GetTaskDefinitionById(definitionId)

	if task.Deleted {
		c.AbortWithError(http.StatusInternalServerError, errors.New("Can't rerun deleted task"))
		return
	}

	poller.DeployJob(task)

	c.JSON(http.StatusOK, gin.H{
		"message": "updated",
	})
}

func GetLatestExecution(c *gin.Context) {

}

func UpdateTaskDefinition(c *gin.Context) {

}

func DeleteTaskDefinition(c *gin.Context) {

}
