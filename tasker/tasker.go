package tasker

import (
	"github.com/drorivry/matter/controllers"
	_ "github.com/drorivry/matter/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	// for _, err := range c.Errors {
	//     log.Fatal("Error", err)
	// }

	// c.JSON(http.StatusInternalServerError, "Internal Server Error")
}

func GetServer() *gin.Engine {
	r := gin.Default()

	r.Use(ErrorHandler)

	r.GET("/ping", controllers.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/task", controllers.GetAllTaskDefinitions)
		v1.POST("/task", controllers.CreateTaskDefinition)
		v1.POST("/task/:definitionId/rerun", controllers.RerunTask)
		v1.GET("/task/:definitionId/latest")
		v1.PUT("/task")
		v1.DELETE("/task/:definitionId")
		v1.GET("/execution", controllers.RerunTask)
		v1.POST("/execution/:executionId/abort", controllers.AbortTaskExecution)
		v1.GET("/tasks/pending", controllers.GetAllPendingTaskDefinitions)
	}

	return r
}
