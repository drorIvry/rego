package tasker

import (
	"net/http"
	"strconv"

	"github.com/drorivry/rego/controllers"
	_ "github.com/drorivry/rego/swagger-docs"
	"github.com/rs/zerolog"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetServer(port int) *http.Server {
	r := gin.Default()
	r.GET("/", controllers.Health)
	r.GET("/ping", controllers.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	ginZerologger := logger.SetLogger(
		logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(gin.DefaultWriter).With().Logger()
		}),
	)

	v1 := r.Group("/api/v1")
	{
		// Definitions
		v1.GET("/task", ginZerologger, controllers.GetAllTaskDefinitions)
		v1.POST("/task", ginZerologger, controllers.CreateTaskDefinition)
		v1.POST("/task/:definitionId/rerun", ginZerologger, controllers.RerunTask)
		v1.GET("/task/:definitionId/latest", ginZerologger, controllers.GetLatestExecution)
		v1.GET("task/:definitionId/history", ginZerologger, controllers.GetTaskHistory)
		v1.PUT("/task", ginZerologger, controllers.UpdateTaskDefinition)
		v1.DELETE("/task/:definitionId", ginZerologger, controllers.DeleteTaskDefinition)
		v1.GET("/tasks/pending", ginZerologger, controllers.GetAllPendingTaskDefinitions)

		// Executions
		v1.GET("/execution", ginZerologger, controllers.RerunTask)
		v1.POST("/execution/:executionId/abort", ginZerologger, controllers.AbortTaskExecution)
	}

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: r,
	}

	return server
}
