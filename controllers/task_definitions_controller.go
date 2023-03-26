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

// CreateTaskDefinition           godoc
// @Summary      				  Create a new task definition
// @Description                   Generate a new definition of a task to run with cadence, parameters and runtime data
// @Tags                          definition
// @Produce                       application/json
// @Param                         newTaskDef  body models.TaskDefinition  true  "Task definition JSON"
// @Success                       200
// @Router                        /task [post]
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

// GetAllTaskDefinitions          godoc
// @Summary      				  Get all task definitions
// @Description                   Filter to get the task definitions you need
// @Tags                          definition
// @Produce                       json
// @Success                       200 {object} []models.TaskDefinition
// @Router                        /task [get]
func GetAllTaskDefinitions(c *gin.Context) {
	tasks := dao.GetAllTaskDefinitions()
	c.IndentedJSON(http.StatusOK, tasks)
}

// GetAllPendingTaskDefinitions   godoc
// @Summary      				  Get all of the pending task definitions
// @Description                   Filter to get the task pending tasks
// @Tags                          definition
// @Produce                       json
// @Success                       200 {object} []models.TaskDefinition
// @Router                        /task/pending [get]
func GetAllPendingTaskDefinitions(c *gin.Context) {
	tasks := dao.GetPendingTasks()
	c.IndentedJSON(http.StatusOK, tasks)
}

// AbortTaskExecution             godoc
// @Summary      				  Rerun a task definition
// @Description                   Rerun a task definition previously created
// @Tags                          definition
// @Produce                       json
// @Param                         definitionId  path string  true  "The task definition id"
// @Success                       200
// @Router                        /task/{definitionId}/rerun [get]
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

// GetLatestExecution	          godoc
// @Summary      				  Get the latest execution of a given definitions
// @Description                   Filter to get the task definitions you need
// @Tags                          definition
// @Produce                       json
// @Success                       200 {object} []models.TaskExecution
// @Param                         definitionId  path string  true  "The task definition id"
// @Router                        /task/{definitionId}/latest [get]
func GetLatestExecution(c *gin.Context) {

}

// UpdateTaskDefinition           godoc
// @Summary      				  Update a task definition
// @Description                   Update a definition of a task to run with cadence, parameters and runtime data
// @Tags                          definition
// @Produce                       application/json
// @Param                         newTaskDef  body models.TaskDefinition  true  "Task definition JSON"
// @Success                       200
// @Router                        /task [put]
func UpdateTaskDefinition(c *gin.Context) {

}

// DeleteTaskDefinition           godoc
// @Summary      				  Delete a task definition
// @Description                   Mark a task definition as deleted (it is not actually deleted from the db)
// @Tags                          definition
// @Produce                       application/json
// @Param                         definitionId  path string  true  "The task definition id"
// @Success                       200
// @Router                        /task/{definitionId} [delete]
func DeleteTaskDefinition(c *gin.Context) {

}
