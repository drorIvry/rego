package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/drorivry/matter/dao"
	"github.com/drorivry/matter/models"
	"github.com/drorivry/matter/poller"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func CreateTaskDefinition(c *gin.Context) {
	var newTaskDef models.TaskDefinition

	if err := c.BindJSON(&newTaskDef); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := dao.CreateTaskDefinition(&newTaskDef)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "created",
	})
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
	definitionId := c.Param("definitionId")
	numericDefinitionId, err := strconv.Atoi(definitionId)
	if err != nil {
		log.Fatal(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	task := dao.GetTaskDefinitionById(uint(numericDefinitionId))

	if task.Deleted {
		c.AbortWithError(http.StatusInternalServerError, errors.New("Can't rerun deleted task"))
		return
	}

	poller.DeployJob(task)

	c.JSON(http.StatusOK, gin.H{
		"message": "updated",
	})
}
