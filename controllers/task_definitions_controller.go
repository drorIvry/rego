package controllers

import (
	"log"
	"net/http"

	"github.com/drorivry/rego/dao"
	"github.com/drorivry/rego/models"
	"github.com/drorivry/rego/poller"
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
// @Router                        /api/v1/task [post]
func CreateTaskDefinition(c *gin.Context) {
	var newTaskDef models.TaskDefinition

	if err := c.BindJSON(&newTaskDef); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	newTaskDef.ID = uuid.New()

	err := dao.CreateTaskDefinition(&newTaskDef)

	if err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(err))
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
// @Router                        /api/v1/task [get]
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
// @Router                        /api/v1/task/pending [get]
func GetAllPendingTaskDefinitions(c *gin.Context) {
	tasks := dao.GetPendingTasks()
	c.IndentedJSON(http.StatusOK, tasks)
}

// RerunTask                      godoc
// @Summary      				  Rerun a task definition
// @Description                   Rerun a task definition previously created
// @Tags                          definition
// @Produce                       json
// @Param                         definitionId  path string  true  "The task definition id"
// @Success                       200
// @Router                        /api/v1/task/{definitionId}/rerun [post]
func RerunTask(c *gin.Context) {
	uuidParam := c.Param("definitionId")
	var definitionId, err = uuid.Parse(uuidParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	task, err := dao.GetTaskDefinitionById(definitionId)

	if task.Deleted {
		c.JSON(http.StatusInternalServerError, c.Error(err))
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
// @Router                        /api/v1/task/{definitionId}/latest [get]
func GetLatestExecution(c *gin.Context) {
	uuidParam := c.Param("definitionId")
	var definitionId, err = uuid.Parse(uuidParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	latestExecution := dao.GetLatestExecutionByDefinitionId(definitionId)
	c.IndentedJSON(http.StatusOK, latestExecution)
}

// UpdateTaskDefinition           godoc
// @Summary      				  Update a task definition
// @Description                   Update a definition of a task to run with cadence, parameters and runtime data
// @Tags                          definition
// @Produce                       application/json
// @Param                         newTaskDef  body models.TaskDefinition  true  "Task definition JSON"
// @Success                       200
// @Router                        /api/v1/task [put]
func UpdateTaskDefinition(c *gin.Context) {
	var updatedTaskDef models.TaskDefinition

	if err := c.BindJSON(&updatedTaskDef); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}
	_, err := dao.GetTaskDefinitionById(updatedTaskDef.ID)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	dao.UpdateDefinition(updatedTaskDef)
	c.JSON(http.StatusOK, gin.H{
		"message": "updated",
	})
}

// DeleteTaskDefinition           godoc
// @Summary      				  Delete a task definition
// @Description                   Mark a task definition as deleted (it is not actually deleted from the db)
// @Tags                          definition
// @Produce                       application/json
// @Param                         definitionId  path string  true  "The task definition id"
// @Success                       200
// @Router                        /api/v1/task/{definitionId} [delete]
func DeleteTaskDefinition(c *gin.Context) {
	uuidParam := c.Param("definitionId")
	var definitionId, err = uuid.Parse(uuidParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}
	dao.DeleteTaskDefinitionById(definitionId)

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}
