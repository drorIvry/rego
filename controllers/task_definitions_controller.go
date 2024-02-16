package controllers

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

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
	apiKey, authErr := AuthRequest(c)

	if authErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var newTaskDef models.TaskDefinition

	if err := c.BindJSON(&newTaskDef); err != nil {
		log.Error().Err(err).Msg("Could not parse task definition")
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	newTaskDef.ID = uuid.New()
	newTaskDef.OrganizationId = apiKey.OrganizationId

	err := dao.CreateTaskDefinition(&newTaskDef)

	if err != nil {
		log.Error().Err(err).Msg("Error creating task definition")
		c.JSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	log.Info().Str("definition_id", newTaskDef.ID.String()).Msg("Created task definition")

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
	apiKey, authErr := AuthRequest(c)

	if authErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tasks := dao.GetAllTaskDefinitions(apiKey.OrganizationId)
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
	apiKey, authErr := AuthRequest(c)

	if authErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	tasks := dao.GetPendingTasks(apiKey.OrganizationId)
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
	apiKey, authErr := AuthRequest(c)

	if authErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	uuidParam := c.Param("definitionId")
	var definitionId, err = uuid.Parse(uuidParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	task, err := dao.GetTaskDefinitionById(definitionId, apiKey.OrganizationId)

	if err == gorm.ErrRecordNotFound {
		log.Warn().Str(
			"definition_id",
			definitionId.String(),
		).Msg("Task definition was not found")
		c.JSON(http.StatusNotFound, c.Error(err))
		return
	} else if err != nil {
		log.Error().Err(err).Msg("Error communicating with the database")
		c.JSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	if task.Deleted {
		c.JSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	poller.DeployJob(&task)

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
	apiKey, authErr := AuthRequest(c)

	if authErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	uuidParam := c.Param("definitionId")
	var definitionId, err = uuid.Parse(uuidParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	latestExecution, err := dao.GetLatestExecutionByDefinitionId(definitionId, apiKey.OrganizationId)

	if err == gorm.ErrRecordNotFound {
		log.Warn().Str(
			"definition_id",
			definitionId.String(),
		).Msg("Task definition was not found")
		c.JSON(http.StatusNotFound, c.Error(err))
		return
	} else if err != nil {
		log.Error().Err(err).Msg("Error communicating with the database")
		c.JSON(http.StatusInternalServerError, c.Error(err))
		return
	}

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
	apiKey, authErr := AuthRequest(c)

	if authErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var updatedTaskDef models.TaskDefinition

	if err := c.BindJSON(&updatedTaskDef); err != nil {
		log.Error().Err(err).Msg("Could not parse updatedTaskDef")
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	updatedTaskDef.OrganizationId = apiKey.OrganizationId
	_, err := dao.GetTaskDefinitionById(updatedTaskDef.ID, apiKey.OrganizationId)

	if err == gorm.ErrRecordNotFound {
		log.Warn().Str(
			"definition_id",
			updatedTaskDef.ID.String(),
		).Msg("Task definition was not found")
		c.JSON(http.StatusNotFound, c.Error(err))
		return
	} else if err != nil {
		log.Error().Err(err).Msg("Error communicating with the database")
		c.JSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	dao.UpdateDefinition(&updatedTaskDef)
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
	apiKey, authErr := AuthRequest(c)

	if authErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	uuidParam := c.Param("definitionId")
	var definitionId, err = uuid.Parse(uuidParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}
	dao.DeleteTaskDefinitionById(definitionId, apiKey.OrganizationId)

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}

// GetTaskHistory                 godoc
// @Summary      				  Returns a history of the executions for the given task definition.
// @Description                   Returns a history of the executions for the given task definition.
// @Tags                          definition
// @Produce                       application/json
// @Param                         definitionId  path  string  true  "The task definition id"
// @Param                         offset  query int  false  "Offset of the history"
// @Param                         limit   query int  false  "Limit the number of the results"
// @Success                       200
// @Router                        /api/v1/task/{definitionId}/history [get]
func GetTaskHistory(c *gin.Context) {
	// var err error
	apiKey, authErr := AuthRequest(c)

	if authErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	uuidParam := c.Param("definitionId")
	definitionId, err := uuid.Parse(uuidParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	offset := ParseIntQueryParameter(c, "offset", 0)
	limit := ParseIntQueryParameter(c, "limit", 10)

	executions, err := dao.GetExecutionsHistory(definitionId, apiKey.OrganizationId, offset, limit)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, executions)
}
