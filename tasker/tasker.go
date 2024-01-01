package tasker

import (
	"net/http"
	"strconv"

	"github.com/drorivry/rego/controllers"
	_ "github.com/drorivry/rego/swagger-docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetServer(port int) *http.Server {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("/api/v1")
	{
		// Definitions
		v1.GET("/task", controllers.GetAllTaskDefinitions)
		v1.POST("/task", controllers.CreateTaskDefinition)
		v1.POST("/task/:definitionId/rerun", controllers.RerunTask)
		v1.GET("/task/:definitionId/latest", controllers.GetLatestExecution)
		v1.PUT("/task", controllers.UpdateTaskDefinition)
		v1.DELETE("/task/:definitionId", controllers.DeleteTaskDefinition)
		v1.GET("/tasks/pending", controllers.GetAllPendingTaskDefinitions)

		// Executions
		v1.GET("/execution", controllers.RerunTask)
		v1.POST("/execution/:executionId/abort", controllers.AbortTaskExecution)
	}

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: r,
	}

	return server
}
